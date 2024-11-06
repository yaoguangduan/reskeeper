package pson

import (
	"encoding/base64"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/yaoguangduan/reskeeper/internal/protox"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Decode(tag string, msg protoreflect.Message, raw string) {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimSuffix(strings.TrimPrefix(raw, "{"), "}")
	fieldKV := splitField(raw)
	for _, kv := range fieldKV {
		idx := strings.Index(kv, ":")
		if idx == -1 {
			log.Panicf("Invalid field key: %s", kv)
		}
		key := strings.TrimSpace(kv[:idx])
		cell := strings.TrimSpace(kv[idx+1:])
		field := msg.Descriptor().Fields().ByName(protoreflect.Name(key))
		if field == nil {
			log.Panicf("message no field key: %s", key)
		}
		if protox.IgnoreCurField(tag, protox.GetMsgTagIgnoreInfo(msg.Descriptor()), field) {
			continue
		}
		if field.IsMap() {
			if !strings.HasPrefix(cell, "{") || !strings.HasSuffix(cell, "}") {
				if !strings.HasPrefix(cell, "[") || !strings.HasSuffix(cell, "]") {
					log.Panicf("Invalid map field: %s", cell)
				}
				cell = strings.TrimSuffix(strings.TrimPrefix(cell, "["), "]")
				mapValList := splitField(cell)
				for _, mv := range mapValList {
					v := ValueOfField(tag, field.MapValue(), mv)
					keyFieldDesc := protox.GetMsgKeyField(field.MapValue().Message())
					if keyFieldDesc == nil {
						panic(fmt.Sprintf("map field must contains key: %s", field))
					}
					k := protoreflect.MapKey(v.Message().Get(*keyFieldDesc))
					msg.Mutable(field).Map().Set(k, v)
				}
			} else {
				cell = strings.TrimSuffix(strings.TrimPrefix(cell, "{"), "}")
				mapKV := splitField(cell)
				for _, mkv := range mapKV {
					idx = strings.Index(mkv, ":")
					if idx == -1 {
						log.Panicf("Invalid map key: %s", mkv)
					}
					mapKey := strings.TrimSpace(mkv[:idx])
					mapCell := strings.TrimSpace(mkv[idx+1:])
					k := ValueOfField(tag, field.MapKey(), mapKey)
					v := ValueOfField(tag, field.MapValue(), mapCell)
					msg.Mutable(field).Map().Set(protoreflect.MapKey(k), v)
				}
			}
		} else {
			var cellList = []string{cell}
			if field.IsList() {
				if !strings.HasPrefix(cell, "[") {
					log.Panicf("Invalid array field: %s", cell)
				}
				cellList = splitField(strings.TrimSuffix(strings.TrimPrefix(cell, "["), "]"))
			}
			for _, val := range cellList {
				value := ValueOfField(tag, field, val)
				if field.IsList() {
					list := msg.Mutable(field).List()
					list.Append(value)
				} else {
					msg.Set(field, value)
				}
			}
		}
	}
}

func ValueOfField(tag string, field protoreflect.FieldDescriptor, cell string) protoreflect.Value {
	var value protoreflect.Value
	switch field.Kind() {
	case protoreflect.StringKind:
		value = protoreflect.ValueOfString(cell)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		u64, err := strconv.ParseUint(cell, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("%s invalid uint32 value: %s,err %e", field.Name(), cell, err))
		}
		if u64 > 4294967295 {
			panic(fmt.Sprintf("cell value %s overflow 4294967295", cell))
		}
		value = protoreflect.ValueOfUint32(uint32(u64))
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		i64, err := strconv.ParseInt(cell, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("%s invalid int32 value: %s,err %e", field.Name(), cell, err))
		}
		if i64 < -2147483648 || i64 > 2147483647 {
			panic(fmt.Sprintf("cell value %s not in -2147483648 and 2147483647", cell))
		}
		value = protoreflect.ValueOfInt32(int32(i64))
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		u64, err := strconv.ParseUint(cell, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("%s invalid uint64 value: %s,err %e", field.Name(), cell, err))
		}
		value = protoreflect.ValueOfUint64(u64)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		i64, err := strconv.ParseInt(cell, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("%s invalid int64 value: %s,err %e", field.Name(), cell, err))
		}
		value = protoreflect.ValueOfInt64(i64)
	case protoreflect.BoolKind:
		value = protoreflect.ValueOfBool(lo.Must(cast.ToBoolE(cell)))
	case protoreflect.DoubleKind:
		f64, err := strconv.ParseFloat(cell, 64)
		if err != nil {
			panic(fmt.Sprintf("%s invalid float64 value: %s,err %e", field.Name(), cell, err))
		}
		value = protoreflect.ValueOfFloat64(f64)
	case protoreflect.FloatKind:
		f64, err := strconv.ParseFloat(cell, 64)
		if err != nil {
			panic(fmt.Sprintf("%s invalid float32 value: %s,err %e", field.Name(), cell, err))
		}
		if f64 < -float64(^uint32(0)>>1)-1 || f64 > float64(^uint32(0)>>1) {
			panic(fmt.Sprintf("cell value %s not in %v and %v", cell, -float64(^uint32(0)>>1)-1, f64 > float64(^uint32(0)>>1)))
		}
		value = protoreflect.ValueOfFloat32(float32(f64))
	case protoreflect.EnumKind:
		value = protoreflect.ValueOfEnum(toEnumNumber(field, cell))
	case protoreflect.BytesKind:
		value = protoreflect.ValueOfBytes(lo.Must(base64.StdEncoding.DecodeString(cell)))
	case protoreflect.MessageKind:
		msg := dynamicpb.NewMessage(field.Message())
		Decode(tag, msg, cell)
		value = protoreflect.ValueOfMessage(msg)
	}
	return value
}

func toEnumNumber(field protoreflect.FieldDescriptor, cell string) protoreflect.EnumNumber {
	evs := field.Enum().Values()
	i32, err := cast.ToInt32E(cell)
	if err == nil {
		return evs.ByNumber(protoreflect.EnumNumber(i32)).Number()
	} else {
		ev := evs.ByName(protoreflect.Name(cell))
		if ev != nil {
			return ev.Number()
		} else {
			eu := protox.GetFieldEnumByAlias(field, cell)
			if eu == -1 {
				panic(fmt.Sprintf("invalid enum value: %s", cell))
			} else {
				return eu
			}
		}
	}
}

func splitField(raw string) []string {
	raw = strings.TrimSpace(raw)
	kv := make([]string, 0)
	var pair = 0
	var beg = 0
	for i := 0; i < len(raw); {
		r, size := rune(raw[i]), 1
		if r >= utf8.RuneSelf {
			r, size = utf8.DecodeRuneInString(raw[i:])
		}
		i += size
		if r == '{' || r == '[' {
			pair += 1
		}
		if r == ']' || r == '}' {
			pair -= 1
		}
		if r == ',' && pair == 0 {
			kv = append(kv, raw[beg:i-1])
			beg = i
		}
	}
	kv = append(kv, raw[beg:])
	return kv
}

func main() {
	//jsonStr := `{age:198,name:a big zoo,pets:[{type:doge,age:12}]}`
	//p := pbgen.User{}
	//zoo := p.ProtoReflect().Descriptor()
	//newZ := dynamicpb.NewMessage(zoo)
	//Decode("", newZ, jsonStr)
	//fmt.Printf("%+v", newZ)
	_, err := cast.ToInt32E(8888888173212)
	if err != nil {
		panic(err)
	}
}

package pson

import (
	"encoding/base64"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/yaoguangduan/reskeeper/internal/protox"
	"github.com/yaoguangduan/reskeeper/pbgen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"log"
	"strings"
	"unicode/utf8"
)

func Decode(msg protoreflect.Message, raw string) {
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
		if field.IsMap() {
			if !strings.HasPrefix(cell, "{") || !strings.HasSuffix(cell, "}") {
				if !strings.HasPrefix(cell, "[") || !strings.HasSuffix(cell, "]") {
					log.Panicf("Invalid map field: %s", cell)
				}
				cell = strings.TrimSuffix(strings.TrimPrefix(cell, "["), "]")
				mapValList := splitField(cell)
				for _, mv := range mapValList {
					v := ValueOfField(field.MapValue(), mv)
					keyFieldDesc := protox.GetMsgKeyField(field.MapValue().Message())
					k := protoreflect.MapKey(v.Message().Get(keyFieldDesc))
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
					k := ValueOfField(field.MapKey(), mapKey)
					v := ValueOfField(field.MapValue(), mapCell)
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
				value := ValueOfField(field, val)
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

func ValueOfField(field protoreflect.FieldDescriptor, cell string) protoreflect.Value {
	var value protoreflect.Value
	switch field.Kind() {
	case protoreflect.StringKind:
		value = protoreflect.ValueOfString(cell)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		value = protoreflect.ValueOfUint32(cast.ToUint32(cell))
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		value = protoreflect.ValueOfInt32(cast.ToInt32(cell))
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		value = protoreflect.ValueOfUint64(cast.ToUint64(cell))
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		value = protoreflect.ValueOfInt64(cast.ToInt64(cell))
	case protoreflect.BoolKind:
		value = protoreflect.ValueOfBool(cast.ToBool(cell))
	case protoreflect.DoubleKind:
		value = protoreflect.ValueOfFloat64(cast.ToFloat64(cell))
	case protoreflect.FloatKind:
		value = protoreflect.ValueOfFloat32(cast.ToFloat32(cell))
	case protoreflect.EnumKind:
		value = protoreflect.ValueOfEnum(toEnumNumber(field, cell))
	case protoreflect.BytesKind:
		value = protoreflect.ValueOfBytes(lo.Must(base64.StdEncoding.DecodeString(cell)))
	case protoreflect.MessageKind:
		msg := dynamicpb.NewMessage(field.Message())
		Decode(msg, cell)
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
		return evs.ByName(protoreflect.Name(cell)).Number()
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
	jsonStr := `{id:198,desc:a big zoo,manager:{age:12,name:makama,assistants:   [{name:AA,level:12},{name:CC,direction:health}]}}`
	p := pbgen.Zoo{}
	zoo := p.ProtoReflect().Descriptor()
	newZ := dynamicpb.NewMessage(zoo)
	Decode(newZ, jsonStr)
	fmt.Printf("%+v", newZ)
}

package protox

import (
	"fmt"
	"github.com/yaoguangduan/reskeeper/resproto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"strconv"
	"strings"
)

func GetFieldByNumber(num int32, msg protoreflect.MessageDescriptor) protoreflect.FieldDescriptor {
	for i := 0; i < msg.Fields().Len(); i++ {
		f := msg.Fields().Get(i)
		if int32(f.Number()) == num {
			return f
		}
	}
	panic(fmt.Sprintf("can not find field by number: %d %s", num, msg))
}

func GetMsgKeyField(msg protoreflect.MessageDescriptor) *protoreflect.FieldDescriptor {
	if proto.HasExtension(msg.Options(), resproto.E_ResMsgKey) {
		fd := msg.Fields().ByNumber(protoreflect.FieldNumber(proto.GetExtension(msg.Options(), resproto.E_ResMsgKey).(int32)))
		return &fd
	}
	return nil
}

func GetFieldByMsgType(fieldMsg protoreflect.MessageDescriptor, msg protoreflect.MessageDescriptor) protoreflect.FieldDescriptor {
	for i := 0; i < msg.Fields().Len(); i++ {
		f := msg.Fields().Get(i)
		if f.Kind() == protoreflect.MessageKind && f.Message() == fieldMsg {
			return f
		}
	}
	panic(fmt.Sprintf("can not find field by number: %d %s", fieldMsg, msg))

}

func GetMsgTagIgnoreInfo(desc protoreflect.MessageDescriptor) []string {
	mo := desc.Options().(*descriptorpb.MessageOptions)
	var msgOpt = make([]string, 0)
	if proto.HasExtension(mo, resproto.E_ResTagIgnores) {
		msgOpt = proto.GetExtension(mo, resproto.E_ResTagIgnores).([]string)
	}
	return msgOpt
}

func IgnoreCurField(tag string, tagIgnoreInfo []string, field protoreflect.FieldDescriptor) bool {
	if len(tagIgnoreInfo) == 0 {
		return false
	}
	for _, tagInfo := range tagIgnoreInfo {
		tagSplit := strings.Split(tagInfo, ":")
		if tagSplit[0] != tag {
			continue
		}
		rangeString, err := parseRangeString(tagSplit[1])
		if err != nil {
			return false
		}
		return isNumberInRange(int(field.Number()), rangeString)
	}
	return false
}

// parseRangeString 解析形如 "数字,数字-数字" 的字符串为数字范围
func parseRangeString(rangeStr string) ([][2]int, error) {
	var ranges [][2]int
	parts := strings.Split(rangeStr, ",")
	for _, part := range parts {
		if strings.Contains(part, "-") {
			bounds := strings.Split(part, "-")
			if len(bounds) != 2 {
				return nil, fmt.Errorf("invalid range format: %s", part)
			}
			start, err1 := strconv.Atoi(bounds[0])
			end, err2 := strconv.Atoi(bounds[1])
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("invalid number in range: %s", part)
			}
			ranges = append(ranges, [2]int{start, end})
		} else {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", part)
			}
			ranges = append(ranges, [2]int{num, num})
		}
	}
	return ranges, nil
}

// isNumberInRange 判断数字是否在给定的范围内
func isNumberInRange(num int, ranges [][2]int) bool {
	for _, r := range ranges {
		if num >= r[0] && num <= r[1] {
			return true
		}
	}
	return false
}

func GetFieldCommentOption(fd protoreflect.FieldDescriptor) string {
	fo := fd.Options()
	if proto.HasExtension(fo, resproto.E_ResComment) {
		return proto.GetExtension(fo, resproto.E_ResComment).(string)
	}
	return ""
}

func GetFieldEnumByAlias(field protoreflect.FieldDescriptor, cell string) protoreflect.EnumNumber {
	values := field.Enum().Values()
	for i := 0; i < values.Len(); i++ {
		ev := values.Get(i)
		evo := ev.Options()
		if proto.HasExtension(evo, resproto.E_ResAlias) && proto.GetExtension(evo, resproto.E_ResAlias).(string) == cell {
			return ev.Number()
		}
	}
	return 0
}

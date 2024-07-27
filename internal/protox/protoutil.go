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

func GetFieldByMsgType(fieldMsg protoreflect.MessageDescriptor, msg protoreflect.MessageDescriptor) protoreflect.FieldDescriptor {
	for i := 0; i < msg.Fields().Len(); i++ {
		f := msg.Fields().Get(i)
		if f.Kind() == protoreflect.MessageKind && f.Message() == fieldMsg {
			return f
		}
	}
	panic(fmt.Sprintf("can not find field by number: %d %s", fieldMsg, msg))

}

func GetMsgOptOrDefault(desc protoreflect.MessageDescriptor) *resproto.ResourceMsgOpt {
	mo := desc.Options().(*descriptorpb.MessageOptions)
	var msgOpt = &resproto.ResourceMsgOpt{TagIgnoreFields: make([]string, 0), MsgKey: proto.String(string(desc.Fields().Get(0).Name())), OneColumn: proto.Bool(false)}
	if proto.HasExtension(mo, resproto.E_ResMsgOpt) {
		msgOpt = proto.GetExtension(mo, resproto.E_ResMsgOpt).(*resproto.ResourceMsgOpt)
	}
	return msgOpt
}

func IgnoreCurField(tag string, opt *resproto.ResourceMsgOpt, field protoreflect.FieldDescriptor) bool {
	if len(opt.TagIgnoreFields) == 0 {
		return false
	}
	for _, tagInfo := range opt.TagIgnoreFields {
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

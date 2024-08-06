package validate

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/yaoguangduan/reskeeper/internal/protox"
	"github.com/yaoguangduan/reskeeper/resproto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"strings"
)

func Validate(tag string, tableMsg protoreflect.Message, tableDesc protoreflect.MessageDescriptor, msgDesc protoreflect.FieldDescriptor) string {
	list := tableMsg.Get(tableDesc.Fields().ByName(msgDesc.Name()))

	buf := &strings.Builder{}
	validateField(buf, tag, list, tableDesc.Fields().ByName(msgDesc.Name()), tableDesc)

	return buf.String()
}

func validateField(buf *strings.Builder, tag string, value protoreflect.Value, field protoreflect.FieldDescriptor, desc protoreflect.MessageDescriptor) {
	var validate *resproto.ResourceFieldValidate
	if !proto.HasExtension(field.Options(), resproto.E_ResValidate) {
		if field.Kind() == protoreflect.MessageKind {
			switch value.Interface().(type) {
			case protoreflect.List:
				list := value.List()
				for i := 0; i < list.Len(); i++ {
					validateMessage(buf, tag, list.Get(i).Message(), field.Message())
				}
			case protoreflect.Map:
				if field.MapValue().Kind() == protoreflect.MessageKind {
					maps := value.Map()
					maps.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
						validateMessage(buf, tag, value.Message(), field.MapValue().Message())
						return true
					})
				}
			default:
				validateMessage(buf, tag, value.Message(), field.Message())
			}
		}
		return
	}
	validate = proto.GetExtension(field.Options(), resproto.E_ResValidate).(*resproto.ResourceFieldValidate)
	if validate.Length != nil {
		beg := cast.ToInt(strings.Split(*validate.Length, "-")[0])
		end := cast.ToInt(strings.Split(*validate.Length, "-")[1])
		switch v := value.Interface().(type) {
		case protoreflect.List:
			if v.Len() < beg || v.Len() >= end {
				buf.WriteString(fmt.Sprintf("Res Validate Error:Length Error msgField:%s[%s] value:%d constraint:%s", desc.Name(), field.Name(), v.Len(), *validate.Length))
			}
		case string:
			if len(v) < beg || len(v) >= end {
				buf.WriteString(fmt.Sprintf("Res Validate Error:Length Error msgField:%s[%s] value:%s constraint:%s", desc.Name(), field.Name(), v, *validate.Length))
			}
		default:
			// skip
			log.Printf("length validate only use for list and string:%s[%s]", desc.Name(), field.Name())
		}
	}
}

func validateMessage(buf *strings.Builder, tag string, value protoreflect.Message, descriptor protoreflect.MessageDescriptor) {
	msgOpt := protox.GetMsgOptOrDefault(descriptor)
	for i := 0; i < descriptor.Fields().Len(); i++ {
		if protox.IgnoreCurField(tag, msgOpt, descriptor.Fields().Get(i)) {
			continue
		}
		validateField(buf, tag, value.Get(descriptor.Fields().Get(i)), descriptor.Fields().Get(i), descriptor)
	}
}

package protox

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
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

package validate

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/yaoguangduan/reskeeper/internal/configs"
	"github.com/yaoguangduan/reskeeper/internal/protox"
	"github.com/yaoguangduan/reskeeper/resproto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"regexp"
	"strings"
	"sync"
)

type validateInfo struct {
	sync.Mutex
	info map[string][]string
}

func (v *validateInfo) Append(ctx configs.CvtContext, field protoreflect.FieldDescriptor, info interface{}) {
	v.Lock()
	defer v.Unlock()
	key := fmt.Sprintf("%s#%s_%s", ctx.Table.GetExcelName(), ctx.Table.GetSheetName(), ctx.Tag)
	_, exist := v.info[key]
	if !exist {
		v.info[key] = make([]string, 0)
	}
	v.info[key] = append(v.info[key], fmt.Sprintf("FIELD %s[%s] validate error: %v\r\n", field.Parent().Name(), field.Name(), info))
}
func (v *validateInfo) String() string {
	sb := strings.Builder{}
	sb.WriteString("RESOURCE VALIDATE RESULT:\r\n")
	return sb.String()
}

var Validator = &validateInfo{info: make(map[string][]string)}

func (v *validateInfo) Validate(tableMsg protoreflect.Message, ctx configs.CvtContext) {
	list := tableMsg.Get(protox.GetFieldByMsgType(ctx.DataDesc, ctx.TableDesc)).List()

	for i := 0; i < list.Len(); i++ {
		v.validateMessage(list.Get(i), ctx.DataDesc, ctx)
	}
}

func (v *validateInfo) validateField(parent protoreflect.Value, field protoreflect.FieldDescriptor, desc protoreflect.MessageDescriptor, ctx configs.CvtContext) {
	if !parent.Message().Has(field) {
		if proto.HasExtension(field.Options(), resproto.E_ResRequire) {
			if proto.GetExtension(field.Options(), resproto.E_ResRequire).(bool) {
				v.Append(ctx, field, "required but empty")
			}
		}
		return
	}
	value := parent.Message().Get(field)
	if proto.HasExtension(field.Options(), resproto.E_ResRange) {
		rge := proto.GetExtension(field.Options(), resproto.E_ResRange).(string)
		beg := cast.ToInt64(strings.Split(rge, "-")[0])
		end := cast.ToUint64(strings.Split(rge, "-")[1])
		switch vv := value.Interface().(type) {
		case protoreflect.List:
			if int64(vv.Len()) < beg || uint64(vv.Len()) > end {
				v.Append(ctx, field, fmt.Sprintf("list size %v not in range %s", vv.Len(), rge))
			}
		case string:
			if int64(len(vv)) < beg || uint64(len(vv)) > end {
				v.Append(ctx, field, fmt.Sprintf("string %s size not in range %s", vv, rge))
			}
		case int32:
			if int64(vv) < beg || uint64(vv) > end {
				v.Append(ctx, field, fmt.Sprintf("%v not in range %s", vv, rge))
			}
		case uint32:
			if int64(vv) < beg || uint64(vv) > end {
				v.Append(ctx, field, fmt.Sprintf("%v not in range %s", vv, rge))
			}
		case uint64:
			if int64(vv) < beg || vv > end {
				v.Append(ctx, field, fmt.Sprintf("%v not in range %s", vv, rge))
			}
		case int64:
			if vv < beg || uint64(vv) > end {
				v.Append(ctx, field, fmt.Sprintf("%v not in range %s", vv, rge))
			}
		default:
			//skip
		}
	}
	if proto.HasExtension(field.Options(), resproto.E_ResPattern) {
		pattern := proto.GetExtension(field.Options(), resproto.E_ResPattern).(string)
		cp, err := regexp.Compile(pattern)
		if err != nil {
			log.Printf("regexp.Compile error:%s %v", pattern, err)
		} else {
			switch vv := value.Interface().(type) {
			case string:
				if !cp.MatchString(vv) {
					v.Append(ctx, field, fmt.Sprintf("string %s not match regex %v", vv, pattern))
				}
			}
		}
	}

	if field.IsList() {
		if proto.HasExtension(field.Options(), resproto.E_ResUnique) && proto.GetExtension(field.Options(), resproto.E_ResUnique).(bool) {
			if field.Kind() != protoreflect.MessageKind {
				for i := 0; i < value.List().Len(); i++ {
					ele := value.List().Get(i)
					for j := i + 1; j < value.List().Len(); j++ {
						eleA := value.List().Get(j)
						if ele.Equal(eleA) {
							v.Append(ctx, field, "need unique but has duplicate element")
						}
					}
				}
			}
		}
	}

	if field.Kind() == protoreflect.MessageKind {
		switch value.Interface().(type) {
		case protoreflect.List:
			list := value.List()
			for i := 0; i < list.Len(); i++ {
				v.validateMessage(list.Get(i), field.Message(), ctx)
			}
		case protoreflect.Map:
			if field.MapValue().Kind() == protoreflect.MessageKind {
				maps := value.Map()
				maps.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
					v.validateMessage(value, field.MapValue().Message(), ctx)
					return true
				})
			}
		default:
			v.validateMessage(value, field.Message(), ctx)
		}
	}
}

func (v *validateInfo) validateMessage(value protoreflect.Value, descriptor protoreflect.MessageDescriptor, ctx configs.CvtContext) {
	msgOpt := protox.GetMsgTagIgnoreInfo(descriptor)
	for i := 0; i < descriptor.Fields().Len(); i++ {
		if protox.IgnoreCurField(ctx.Tag, msgOpt, descriptor.Fields().Get(i)) {
			continue
		}
		v.validateField(value, descriptor.Fields().Get(i), descriptor, ctx)
	}
}

func (v *validateInfo) PrintValidateResult() {
	if len(v.info) == 0 {
		return
	}
	sb := strings.Builder{}

	sb.WriteString("RESOURCE VALIDATE RESULT:\r\n")
	for key, val := range v.info {
		if len(val) > 0 {
			sb.WriteString(key + ":\r\n")
			for _, vv := range val {
				sb.WriteString("  " + vv)
			}
		} else {
			sb.WriteString(key + " validate ok")
		}
	}
	log.Println(sb.String())
}

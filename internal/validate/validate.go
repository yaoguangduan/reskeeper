package validate

//
//import (
//	"fmt"
//	"github.com/spf13/cast"
//	"github.com/yaoguangduan/reskeeper/internal/configs"
//	"github.com/yaoguangduan/reskeeper/internal/protox"
//	"github.com/yaoguangduan/reskeeper/resproto"
//	"google.golang.org/protobuf/proto"
//	"google.golang.org/protobuf/reflect/protoreflect"
//	"log"
//	"regexp"
//	"strings"
//	"sync"
//)
//
//type validateInfo struct {
//	sync.Mutex
//	info map[string][]string
//}
//
//func (v *validateInfo) Append(ctx configs.CvtContext, typ string, field protoreflect.FieldDescriptor, real, expect interface{}) {
//	v.Lock()
//	defer v.Unlock()
//	key := fmt.Sprintf("%s#%s_%s", ctx.Table.GetExcelName(), ctx.Table.GetSheetName(), ctx.Tag)
//	_, exist := v.info[key]
//	if !exist {
//		v.info[key] = make([]string, 0)
//	}
//	v.info[key] = append(v.info[key], fmt.Sprintf("VALIDATE ERROR TYPE:%s FIELD:%s[%s] REAL:%v EXPECT:%v\r\n", typ, field.Parent().Name(), field.Name(), real, expect))
//}
//func (v *validateInfo) String() string {
//	sb := strings.Builder{}
//	sb.WriteString("RESOURCE VALIDATE RESULT:\r\n")
//	return sb.String()
//}
//
//var Validator = &validateInfo{info: make(map[string][]string)}
//
//func (v *validateInfo) Validate(tableMsg protoreflect.Message, ctx configs.CvtContext) {
//	list := tableMsg.Get(protox.GetFieldByMsgType(ctx.DataDesc, ctx.TableDesc)).List()
//
//	for i := 0; i < list.Len(); i++ {
//		v.validateMessage(list.Get(i), ctx.DataDesc, ctx)
//	}
//}
//
//func (v *validateInfo) validateField(parent protoreflect.Value, field protoreflect.FieldDescriptor, desc protoreflect.MessageDescriptor, ctx configs.CvtContext) {
//	var validate *resproto.ResourceFieldValidate
//	if !parent.Message().Has(field) {
//		if proto.HasExtension(field.Options(), resproto.E_ResValidate) {
//			validate = proto.GetExtension(field.Options(), resproto.E_ResValidate).(*resproto.ResourceFieldValidate)
//			if validate.NotNull != nil && *validate.NotNull {
//				v.Append(ctx, "NotNull", field, "", "")
//			}
//		}
//		return
//	}
//	if proto.HasExtension(field.Options(), resproto.E_ResValidate) {
//		validate = proto.GetExtension(field.Options(), resproto.E_ResValidate).(*resproto.ResourceFieldValidate)
//	}
//
//	if !proto.HasExtension(field.Options(), resproto.E_ResValidate) {
//		if field.Kind() == protoreflect.MessageKind {
//			value := parent.Message().Get(field)
//			switch value.Interface().(type) {
//			case protoreflect.List:
//				list := value.List()
//				for i := 0; i < list.Len(); i++ {
//					v.validateMessage(list.Get(i), field.Message(), ctx)
//				}
//			case protoreflect.Map:
//				if field.MapValue().Kind() == protoreflect.MessageKind {
//					maps := value.Map()
//					maps.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
//						v.validateMessage(value, field.MapValue().Message(), ctx)
//						return true
//					})
//				}
//			default:
//				v.validateMessage(value, field.Message(), ctx)
//			}
//		}
//		return
//	}
//	value := parent.Message().Get(field)
//	if validate.Length != nil {
//		beg := cast.ToInt(strings.Split(*validate.Length, "-")[0])
//		end := cast.ToInt(strings.Split(*validate.Length, "-")[1])
//		switch vv := value.Interface().(type) {
//		case protoreflect.List:
//			if vv.Len() < beg || vv.Len() >= end {
//				v.Append(ctx, "Length", field, vv.Len(), *validate.Length)
//			}
//		case string:
//			if len(vv) < beg || len(vv) >= end {
//				v.Append(ctx, "Length", field, len(vv), *validate.Length)
//			}
//		default:
//			// skip
//			log.Printf("length validate only use for list and string:%s[%s]", desc.Name(), field.Name())
//		}
//	}
//	if validate.Range != nil {
//		beg := cast.ToInt64(strings.Split(*validate.Range, "-")[0])
//		end := cast.ToUint64(strings.Split(*validate.Range, "-")[1])
//		switch vv := value.Interface().(type) {
//		case int32:
//			if int64(vv) < beg || uint64(vv) >= end {
//				v.Append(ctx, "Range", field, vv, *validate.Range)
//			}
//		case uint32:
//			if int64(vv) < beg || uint64(vv) >= end {
//				v.Append(ctx, "Range", field, vv, *validate.Range)
//			}
//		case uint64:
//			if int64(vv) < beg || vv >= end {
//				v.Append(ctx, "Range", field, vv, *validate.Range)
//			}
//		case int64:
//			if vv < beg || uint64(vv) >= end {
//				v.Append(ctx, "Range", field, vv, *validate.Range)
//			}
//		default:
//			//skip
//		}
//	}
//	if validate.Pattern != nil {
//		cp, err := regexp.Compile(*validate.Pattern)
//		if err != nil {
//			log.Printf("regexp.Compile error:%s %v", *validate.Pattern, err)
//		} else {
//			switch vv := value.Interface().(type) {
//			case string:
//				if !cp.MatchString(vv) {
//					v.Append(ctx, "Pattern", field, vv, *validate.Pattern)
//				}
//			}
//		}
//	}
//}
//
//func (v *validateInfo) validateMessage(value protoreflect.Value, descriptor protoreflect.MessageDescriptor, ctx configs.CvtContext) {
//	//msgOpt := protox.GetMsgOptOrDefault(descriptor)
//	//for i := 0; i < descriptor.Fields().Len(); i++ {
//	//	if protox.IgnoreCurField(ctx.Tag, msgOpt, descriptor.Fields().Get(i)) {
//	//		continue
//	//	}
//	//	v.validateField(value, descriptor.Fields().Get(i), descriptor, ctx)
//	//}
//}
//
//func (v *validateInfo) PrintValidateResult() {
//	sb := strings.Builder{}
//	sb.WriteString("=======================================\r\n")
//	sb.WriteString("RESOURCE VALIDATE RESULT:\r\n")
//	for key, val := range v.info {
//		sb.WriteString(key + ":\r\n")
//		for _, vv := range val {
//			sb.WriteString("  " + vv)
//		}
//	}
//	sb.WriteString("=======================================\r\n")
//	fmt.Println(sb.String())
//}

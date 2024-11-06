package protox

import (
	"fmt"
	"github.com/samber/lo"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"unicode"
)

type ProtoFiles struct {
	RegFiles *protoregistry.Files
}

func (p ProtoFiles) GetMessage(msgName string) protoreflect.MessageDescriptor {
	var md protoreflect.MessageDescriptor
	p.RegFiles.RangeFiles(func(descriptor protoreflect.FileDescriptor) bool {
		for i := 0; i < descriptor.Messages().Len(); i++ {
			if string(descriptor.Messages().Get(i).Name()) == msgName {
				md = descriptor.Messages().Get(i)
				return false
			}
		}
		return true
	})
	if md == nil {
		log.Fatalf("cannot find msg:%s", msgName)
	}
	return md
}

func ParseProtoSet(config []string) ProtoFiles {
	name := protocGen(config)
	pf := readProtoDescSet(name)
	return pf
}

var keyValidTypes = []protoreflect.Kind{protoreflect.BoolKind, protoreflect.EnumKind, protoreflect.Int32Kind, protoreflect.Uint32Kind, protoreflect.Sint32Kind, protoreflect.Int64Kind, protoreflect.Uint64Kind, protoreflect.Sint64Kind, protoreflect.StringKind}

func readProtoDescSet(name string) ProtoFiles {
	defer func() {
		lo.Must0(os.Remove(name))
	}()
	//读取描述符集文件
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("Failed to read descriptor set: %v", err)
	}

	var fileSet descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(data, &fileSet); err != nil {
		log.Fatalf("Failed to unmarshal descriptor set: %v", err)
	}
	files, err := protodesc.NewFiles(&fileSet)
	if err != nil {
		log.Fatalf("Failed to create descriptor set: %v", err)
	}
	return ProtoFiles{RegFiles: files}
}

func protocGen(protoDirs []string) string {
	seconds := time.Now().Unix()
	dsName := fmt.Sprintf("desc_set_%d.pb", seconds)
	arg1List := lo.Map(protoDirs, func(item string, index int) string {
		return fmt.Sprintf("--proto_path=%s", filepath.ToSlash(filepath.Clean(item)))
	})
	arg1List = append(arg1List, "--proto_path=.")
	arg1List = append(arg1List, "--proto_path=./google/protobuf")
	arg2List := lo.Map(protoDirs, func(item string, index int) string {
		return filepath.ToSlash(filepath.Join(filepath.ToSlash(filepath.Clean(item)), "*.proto"))
	})
	arg2List = append(arg2List, "resource_opt.proto")
	arg2List = append(arg2List, "./google/protobuf/descriptor.proto")
	cmd := exec.Command("./protoc.exe", append([]string{"--descriptor_set_out=" + dsName}, append(arg1List, arg2List...)...)...)
	fmt.Println(cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("protoc.exe failed: %v\n%s", err, out))
	}
	return dsName
}

// ToLowerFirst 将字符串的首字母转换为小写
func ToLowerFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

package protox

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reskeeper/res_toml"
	"strings"
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

func ParseProtoSet(config res_toml.Config) ProtoFiles {
	name := protocGen(config)
	defer func() {
		lo.Must0(os.Remove(name))
	}()
	pf := readProtoDescSet(name, config)
	checkProtosValid(config, pf)
	return pf
}

var keyValidTypes = []protoreflect.Kind{protoreflect.BoolKind, protoreflect.EnumKind, protoreflect.Int32Kind, protoreflect.Uint32Kind, protoreflect.Sint32Kind, protoreflect.Int64Kind, protoreflect.Uint64Kind, protoreflect.Sint64Kind, protoreflect.StringKind}

func checkProtosValid(config res_toml.Config, pf ProtoFiles) {
	for _, table := range config.Tables {
		msg := pf.GetMessage(table.MessageName)
		field := GetFieldByNumber(1, msg)
		if !lo.Contains(keyValidTypes, field.Kind()) {
			panic(fmt.Sprintf("msg invalid key %s[%s] %s", field.Name(), field.Kind(), msg.Name()))
		}
	}
}

func readProtoDescSet(name string, config res_toml.Config) ProtoFiles {
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

func protocGen(config res_toml.Config) string {
	seconds := time.Now().Unix()
	dsName := fmt.Sprintf("desc_set_%d.pb", seconds)
	cmd := exec.Command("./protoc.exe", "--descriptor_set_out="+dsName, "--proto_path="+config.ProtoPath, config.ProtoPath+"/*proto")
	fmt.Println(cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		panic(err)
	}
	fmt.Println("Protobuf file set gen successfully.")
	return dsName
}

func getOneGoPackageLine(config res_toml.Config) string {
	pp := config.ProtoPath

	entries, err := os.ReadDir(pp)
	if err != nil {
		log.Fatalf("Error reading directory %q: %v\n", pp, err)
	}

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".proto") {
			fr := lo.Must(os.Open(filepath.Join(pp, entry.Name())))
			scan := bufio.NewScanner(fr)
			for scan.Scan() {
				line := scan.Text()
				if strings.Contains(line, "option") && strings.Contains(line, "go_package") {
					return line
				}
			}
		}
	}
	return ""
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

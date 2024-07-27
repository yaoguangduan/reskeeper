package main

import (
	"flag"
	"github.com/yaoguangduan/reskeeper/internal"
	_ "google.golang.org/protobuf/reflect/protodesc"
	"strings"
)

// 定义一个自定义标志类型
type protoList []string

func (s *protoList) String() string {
	return strings.Join(*s, ",")
}

func (s *protoList) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	var list protoList
	flag.Var(&list, "list", "a list of strings")
	flag.Parse()

	internal.Gen(list)

}

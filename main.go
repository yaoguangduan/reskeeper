package main

import (
	"flag"
	"github.com/yaoguangduan/reskeeper/internal"
	_ "google.golang.org/protobuf/reflect/protodesc"
	"strings"
)

type StringList []string

func (s *StringList) String() string {
	return strings.Join(*s, ",")
}

func (s *StringList) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	var protoDirs StringList
	var marshalExcelSheetList StringList
	flag.Var(&protoDirs, "P", "a list of proto dir")
	genExcelAndSheet := flag.Bool("E", true, "generate missing excel files and sheets,default is true")
	flag.Var(&marshalExcelSheetList, "C", "a list of excel name or excel#sheet")
	flag.Parse()

	internal.Gen(protoDirs, marshalExcelSheetList, *genExcelAndSheet)

}

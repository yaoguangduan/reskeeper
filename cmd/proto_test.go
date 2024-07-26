package main

import (
	"fmt"
	"github.com/yaoguangduan/reskeeper/pbgen"
	"github.com/yaoguangduan/reskeeper/resproto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"testing"
)

func TestGene(t *testing.T) {
	//file, err := os.ReadFile("..\\data\\Zoo_full.bin")
	//if err != nil {
	//	panic(err)
	//}
	zooTable := pbgen.ZooTable{}
	//err = proto.Unmarshal(file, &zooTable)
	//for _, zoo := range zooTable.Zoos {
	//	for _, b := range zoo.Borrows {
	//		fmt.Println(b)
	//	}
	//}
	//fmt.Println(&zooTable)
	desc := zooTable.ProtoReflect().Descriptor()
	mo := desc.Options().(*descriptorpb.MessageOptions)
	e := proto.GetExtension(mo, resproto.E_ExcelAndSheetName)
	fmt.Println(e)
}

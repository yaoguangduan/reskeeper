package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"os"
	"reskeeper/pbgen"
	"testing"
)

func TestGene(t *testing.T) {
	file, err := os.ReadFile("..\\data\\Zoo_full.bin")
	if err != nil {
		panic(err)
	}
	zooTable := pbgen.ZooTable{}
	err = proto.Unmarshal(file, &zooTable)
	for _, zoo := range zooTable.Zoos {
		for _, b := range zoo.Borrows {
			fmt.Println(b)
		}
	}
	fmt.Println(&zooTable)
}

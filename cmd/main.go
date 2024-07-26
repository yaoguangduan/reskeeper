package main

import (
	"github.com/yaoguangduan/reskeeper/internal"
	_ "google.golang.org/protobuf/reflect/protodesc"
	"os"
)

func main() {

	var file = "resource.toml"
	if len(os.Args) >= 2 {
		file = os.Args[1]
	}
	internal.Gen(file)

}

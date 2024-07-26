package main

import (
	_ "google.golang.org/protobuf/reflect/protodesc"
	"os"
	"reskeeper/internal"
)

func main() {

	var file = "resource.toml"
	if len(os.Args) >= 2 {
		file = os.Args[1]
	}
	internal.Gen(file)

}

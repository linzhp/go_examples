package main

import (
	"go/build"
	"fmt"
)

func main() {
	build.Default.BuildTags = []string{"freebsd"}
	pkg, err := build.Import("./unix", "./testdata", 0)
	if err == nil {
		fmt.Println(pkg.GoFiles)
	} else {
		fmt.Println(err)
	}
}

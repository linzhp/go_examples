package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

var testFileSet = token.NewFileSet()

func main() {
	filename := "power-driven-gen.go"
	_, err := parser.ParseFile(testFileSet, filename, nil, parser.ParseComments)
	fmt.Println(err)
}

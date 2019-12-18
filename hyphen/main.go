package main

import (
	"fmt"
	"go/parser"
	"go/token"

	_ "github.com/StackExchange/dnscontrol/providers/hexonet"
	_ "github.com/linzhp/go_examples/modEnabled/subPkg"
)

var testFileSet = token.NewFileSet()

func main() {
	filename := "power-driven-gen.go"
	_, err := parser.ParseFile(testFileSet, filename, nil, parser.ParseComments)
	fmt.Println(err)
}

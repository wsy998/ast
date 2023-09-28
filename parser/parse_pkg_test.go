package parser_test

import (
	"fmt"
	"testing"

	"github.com/wsy998/ast/v1/parser"
)

func TestParseFile(t *testing.T) {
	parse, err := parser.Parse("testdata/a.go")
	if err != nil {
		return
	}
	fmt.Println(parse)
}

func TestParseFile2(t *testing.T) {
	parse, err := parser.ParsePackage("testdata")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parse)
}

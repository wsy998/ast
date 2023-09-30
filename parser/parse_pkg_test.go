package parser_test

import (
	"fmt"
	"testing"

	"github.com/wsy998/ast/parser"
)

func TestParseFile(t *testing.T) {
	parse, err := parser.Parse("testdata/b.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parse)
}

func TestParseFile2(t *testing.T) {
	parse, err := parser.Parse("testdata/c.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parse)
}
func TestParseFile3(t *testing.T) {
	parse, err := parser.Parse("testdata/c.go")
	if err != nil {
		return
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parse)
}

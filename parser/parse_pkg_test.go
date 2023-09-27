package parser_test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/wsy998/ast/parser"
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
		return
	}
	fmt.Println(parse.OpenReceiver())
	fmt.Println(parse.WithoutReceiver())
	fmt.Println(parse.OpenWithoutReceiver())
}
func TestS(t *testing.T) {
	sizeof := unsafe.Sizeof(struct {
		Open        bool
		Field       []*parser.GoField
		Funs        []*parser.GoFunc
		Package     map[string]string
		Constructor *parser.GoFunc
		Comment     string
		Name        string
	}{})
	fmt.Println(sizeof)
}

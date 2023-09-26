package parser

import (
	"os"
	"strings"
)

func ParsePackage(pkg string) error {
	dir, err := os.ReadDir(pkg)
	if err != nil {
		return nil
	}
	structs := make([]*GoStruct, 0)
	funcs := make([]*GoFunc, 0)
	for _, entry := range dir {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}
		parse, err := Parse(entry.Name())
		if err != nil {
			return err
		}
		for _, goFunc := range parse.Func {
			funcs = append(funcs, goFunc)
		}
		for _, goFunc := range parse.GoStructs {
			structs = append(structs, goFunc)
		}
	}
	goPkg := NewGoPkg()
	goPkg.GoStructs = structs
	goPkg.Func = funcs
	return nil
}

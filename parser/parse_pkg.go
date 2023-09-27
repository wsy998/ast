package parser

import (
	"os"
	"path/filepath"
	"strings"
)

func ParsePackage(pkg string) (*GoPkg, error) {
	dir, err := os.ReadDir(pkg)
	if err != nil {
		return nil, err
	}
	structs := make([]*GoStruct, 0)
	funcs := make([]*GoFunc, 0)
	imports := make(map[string]string)
	for _, entry := range dir {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}
		parse, err := Parse(filepath.Join(pkg, entry.Name()))
		if err != nil {
			return nil, err
		}
		for _, goFunc := range parse.Func {
			funcs = append(funcs, goFunc)
		}
		for _, goFunc := range parse.GoStructs {
			structs = append(structs, goFunc)
		}
		for n, goImport := range parse.Imports {
			imports[n] = goImport
		}
	}
	goPkg := NewGoPkg()
	goPkg.GoStructs = structs
	goPkg.Func = funcs
	goPkg.Imports = imports
	return goPkg, nil
}

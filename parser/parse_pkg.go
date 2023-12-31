package parser

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/wsy998/ast/api"
)

func ParsePackage(pkg string) (*api.GoPkg, error) {
	dir, err := os.ReadDir(pkg)
	if err != nil {
		return nil, err
	}
	structs := make([]*api.GoStruct, 0)
	funcs := make([]*api.GoFunc, 0)
	imports := make(map[string]string)
	for _, entry := range dir {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}
		set := token.NewFileSet()
		file, err := parser.ParseFile(set, filepath.Join(pkg, entry.Name()), nil, 0)
		if err != nil {
			return nil, err
		}
		parse, err := parse(file, true)
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

	for _, goFunc := range funcs {
		for _, goStruct := range structs {
			if isConstructor(goFunc, goStruct) {
				goStruct.Constructor = goFunc
				break
			}
			if IsMethod(goFunc, goStruct) {
				goStruct.Funs = append(goStruct.Funs, goFunc)
			}
		}
	}
	goPkg := api.NewGoPkg()
	goPkg.GoStructs = structs
	goPkg.Func = funcs
	goPkg.Imports = imports
	goPkg.Name = filepath.Base(pkg)
	return goPkg, nil
}

func ParsePackages(pkgs map[string]*api.GoPkg, re bool, pkg ...string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	goPkg := api.NewGoPkg()
	for _, s := range pkg {
		dir, err := os.ReadDir(filepath.Join(pwd, s))
		if err != nil {
			return err
		}
		for _, entry := range dir {
			if entry.IsDir() {
				if re {
					err := ParsePackages(pkgs, re, s+"/"+entry.Name())
					if err != nil {
						return err
					}
					continue
				}
			}
			if !strings.HasSuffix(entry.Name(), ".go") {
				continue
			}

			file, err := Parse(s + "/" + entry.Name())
			if err != nil {
				return err
			}
			goPkg.Name = file.Pkg
			goPkg.GoStructs = append(goPkg.GoStructs, file.GoStructs...)
			goPkg.Func = append(goPkg.Func, file.Func...)
			for name, imports := range file.Imports {
				goPkg.Imports[name] = imports
			}
			for _, goFunc := range goPkg.Func {
				for _, goStruct := range goPkg.GoStructs {
					if isConstructor(goFunc, goStruct) {
						goStruct.Constructor = goFunc
						break
					}
					if IsMethod(goFunc, goStruct) {
						goStruct.Funs = append(goStruct.Funs, goFunc)
					}
				}
			}
			pkgs[s] = goPkg

		}

	}

	return nil
}

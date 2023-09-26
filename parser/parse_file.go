package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func Parse(file string) (*GoFile, error) {
	goFile := NewGoFile()
	set := token.NewFileSet()
	parseFile, err := parser.ParseFile(set, file, nil, 0)
	if err != nil {
		return nil, err
	}
	ast.Print(set, parseFile)
	decls := parseFile.Decls
	for _, decl := range decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if goFile.Func == nil {
				goFile.Func = make([]*GoFunc, 0)
			}

			goFunc := NewGoFunc()
			field := parseField(d.Recv)
			if len(field) > 0 {
				goFunc.Receiver = field[0]
			}
			goFunc.Open = d.Name.IsExported()
			goFunc.Name = d.Name.Name
			goFunc.Comment = d.Doc.Text()
			goFunc.In = parseField(d.Type.Params)
			goFunc.Out = parseField(d.Type.Results)
			goFile.Func = append(goFile.Func, goFunc)
		case *ast.GenDecl:
			if d.Tok == token.TYPE {
				for _, spec := range d.Specs {
					if v, ok := spec.(*ast.TypeSpec); ok {
						if s, o := v.Type.(*ast.StructType); o {
							goStruct := NewGoStruct()
							goStruct.Name = v.Name.Name
							goStruct.Comment = v.Comment.Text()
							goStruct.Field = parseField(s.Fields)
							goStruct.Open = v.Name.IsExported()
							goFile.GoStructs = append(goFile.GoStructs, goStruct)
						}
					}
				}
			}
		}
	}

	for _, goFunc := range goFile.Func {
		for _, goStruct := range goFile.GoStructs {
			if goFunc.Receiver.Type == goStruct.Name {
				if goStruct.Fun == nil {
					goStruct.Fun = make([]*GoFunc, 0)
				}
				goStruct.Fun = append(goStruct.Fun, goFunc)
			}
		}

	}

	return goFile, nil
}

func parseField(field *ast.FieldList) []*GoField {
	if field.NumFields() == 0 {
		return nil
	}
	k := make([]*GoField, 0)
	list := field.List
	for _, f := range list {
		for _, name := range f.Names {
			goField := NewGoField()
			goField.Open = name.IsExported()
			goField.Name = name.Name
			goField.Type, goField.Pointer = parseFieldType(f.Type)
			goField.Tag = make(map[string]string)
			if f.Tag != nil && len(f.Tag.Value) > 0 {
				tag := unWrapTag(f.Tag.Value)
				for _, v := range strings.Split(tag, " ") {
					v = strings.Trim(v, " ")
					if v != "" {
						indexByte := strings.IndexByte(v, ':')
						name := v[:indexByte]
						value := unWrapTag(v[indexByte+1:])
						goField.Tag[name] = value
					}
				}
			}
			k = append(k, goField)
		}
	}
	return k
}

func unWrapTag(str string) string {
	s := []string{"`", `"`}
	sl := ""
	for _, s2 := range s {
		if strings.HasPrefix(str, s2) && strings.HasSuffix(str, s2) {
			sl = str[1 : len(str)-1]
		}

	}
	return sl
}

func parseFieldType(expr ast.Expr) (string, bool) {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name, false
	case *ast.SelectorExpr:
		return e.X.(*ast.Ident).Name + "." + e.Sel.Name, false
	case *ast.StarExpr:
		switch rt := e.X.(type) {
		case *ast.Ident:
			return rt.Name, true
		case *ast.SelectorExpr:
			return rt.X.(*ast.Ident).Name + "." + rt.Sel.Name, true
		}

	}
	return "", false
}

package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/wsy998/ast/consts"
	"github.com/wsy998/ast/util"
)

func Parse(file string) (*GoFile, error) {
	goFile := NewGoFile()
	set := token.NewFileSet()
	parseFile, err := parser.ParseFile(set, file, nil, 0)
	if err != nil {
		return nil, err
	}
	importMap := make(map[string]string)
	for _, spec := range parseFile.Imports {
		value := util.UnwrapQuote(spec.Path.Value)
		index := strings.LastIndexByte(value, consts.Slash)
		name := value[index+1:]
		if spec.Name != nil && len(spec.Name.Name) > 0 {
			name = util.UnwrapQuote(spec.Name.Name)
		}
		importMap[name] = value
	}
	goFile.Imports = importMap
	decls := parseFile.Decls
	for _, decl := range decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if goFile.Func == nil {
				goFile.Func = make([]*GoFunc, 0)
			}

			goFunc := NewGoFunc()
			field := parseField(d.Recv, importMap)
			if len(field) > 0 {
				goFunc.Receiver = field[0]
			}
			goFunc.Open = d.Name.IsExported()
			goFunc.Name = d.Name.Name
			goFunc.Comment = d.Doc.Text()
			goFunc.In = parseField(d.Type.Params, importMap)
			goFunc.Out = parseField(d.Type.Results, importMap)
			goFile.Func = append(goFile.Func, goFunc)
		case *ast.GenDecl:
			if d.Tok == token.TYPE {
				for _, spec := range d.Specs {
					if v, ok := spec.(*ast.TypeSpec); ok {
						if s, o := v.Type.(*ast.StructType); o {
							goStruct := NewGoStruct()
							goStruct.Name = v.Name.Name
							goStruct.Comment = v.Comment.Text()
							goStruct.Field = parseField(s.Fields, importMap)
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
			if goFunc.Name == "New"+util.UcFirst(goStruct.Name) {
				if goFunc.Receiver != nil || len(goFunc.In) != 0 || len(goFunc.Out) != 1 {
					continue
				}
				if goFunc.Out[0].Type == goStruct.Name {
					goStruct.Constructor = goFunc
				}
			}
		}

	}

	return goFile, nil
}

func parseField(field *ast.FieldList, importMap map[string]string) []*GoField {
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
			goField.Package, goField.Type, goField.Pointer, goField.Field = parseFieldType(f.Type, importMap)
			goField.Tag = make(map[string]string)
			if f.Tag != nil && len(f.Tag.Value) > 0 {
				tag := util.UnwrapQuote(f.Tag.Value)
				for _, v := range strings.Split(tag, string(consts.Space)) {
					v = strings.Trim(v, string(consts.Space))
					if !util.EmptyString(v) {
						indexByte := strings.IndexByte(v, consts.Colon)
						name := v[:indexByte]
						value := util.UnwrapQuote(v[indexByte+1:])
						goField.Tag[name] = value
					}
				}
			}
			k = append(k, goField)
		}
	}
	return k
}

func parseFieldType(expr ast.Expr, importMap map[string]string) (string, string, bool, []*GoField) {

	switch e := expr.(type) {
	case *ast.Ident:
		return consts.Empty, e.Name, false, nil
	case *ast.SelectorExpr:
		return e.X.(*ast.Ident).Name, fmt.Sprintf("%s.%s", e.X.(*ast.Ident).Name, e.Sel.Name), false, nil
	case *ast.StructType:
		field := parseField(e.Fields, importMap)
		return consts.Empty, consts.TypeStruct, false, field
	case *ast.StarExpr:
		switch rt := e.X.(type) {
		case *ast.Ident:
			return consts.Empty, rt.Name, true, nil
		case *ast.SelectorExpr:
			return rt.X.(*ast.Ident).Name, fmt.Sprintf("%s.%s", rt.X.(*ast.Ident).Name, rt.Sel.Name), true, nil
		case *ast.StructType:
			field := parseField(rt.Fields, importMap)
			return consts.Empty, consts.TypeStruct, true, field
		}

	}
	return consts.Empty, consts.Empty, false, nil
}

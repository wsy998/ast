package parser

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"strings"

	"github.com/wsy998/ast/api"
	const2 "github.com/wsy998/ast/consts"
	"github.com/wsy998/ast/internal/consts"
	"github.com/wsy998/ast/internal/util"
)

func Parse(file string) (*api.GoFile, error) {
	fileP, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fileP.Close()
	return ParseReader(fileP)
}

func ParseReader(reader io.Reader) (*api.GoFile, error) {
	buffer := bytes.NewBuffer(nil)
	_, err := io.Copy(buffer, reader)
	if err != nil && err != io.EOF {
		return nil, err
	}
	set := token.NewFileSet()
	parseFile, err := parser.ParseFile(set, "", buffer.Bytes(), 0)
	if err != nil {
		return nil, err
	}

	return parse(parseFile, false)
}

func parse(parseFile *ast.File, delay bool) (*api.GoFile, error) {

	m := parseImport(parseFile)

	file := api.NewGoFile()
	file.Pkg = parseFile.Name.Name
	if parseFile.Decls != nil && len(parseFile.Decls) > 0 {
		for _, decl := range parseFile.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				genDecl := parseGenDecl(d, m)
				if genDecl != nil {
					file.GoStructs = append(file.GoStructs, genDecl)
				}
			case *ast.FuncDecl:
				file.Func = append(file.Func, parseFuncDecl(d, m))
			}

		}
	}
	if !delay {
		for _, goFunc := range file.Func {
			for _, goStruct := range file.GoStructs {
				if isConstructor(goFunc, goStruct) {
					goStruct.Constructor = goFunc
					break
				}
				if IsMethod(goFunc, goStruct) {
					goStruct.Funs = append(goStruct.Funs, goFunc)
				}
			}

		}
	}
	return file, nil
}

func IsMethod(goFunc *api.GoFunc, goStruct *api.GoStruct) bool {
	if goFunc.Receiver == nil {
		return false
	}
	if goFunc.Receiver.Type.TypeName() == const2.Pointer {
		rt := goFunc.Receiver.Type.RealType()
		if rt.TypeName() == "ident" {
			if goStruct.Name == rt.String() {
				return true
			}
		}
	} else if goFunc.Receiver.Type.TypeName() == "ident" {
		if goStruct.Name == goFunc.Receiver.Type.String() {
			return true
		}
	}
	return false
}

func isConstructor(goFunc *api.GoFunc, goStruct *api.GoStruct) bool {
	if goFunc.Receiver != nil {
		return false
	}
	if goFunc.Name != "New"+goStruct.Name {
		return false
	}
	if goFunc.Out == nil || len(goFunc.Out) == 0 {
		return false
	}
	if len(goFunc.Out) > 0 {
		if goFunc.Out[0].Type.TypeName() == const2.Pointer {
			rt := goFunc.Out[0].Type.RealType()
			if rt.String() == goStruct.Name {
				return true
			}
		} else if goFunc.Out[0].Type.TypeName() == "ident" {
			if goFunc.Out[0].Type.String() == goStruct.Name {
				return true
			}
		} else {
			return false
		}

	}
	return false
}

func parseFuncDecl(d *ast.FuncDecl, m map[string]string) *api.GoFunc {
	funcType := api.NewGoFunc()
	funcType.Name = d.Name.Name
	funcType.In = parseField(d.Type.Params, m)
	funcType.Out = parseField(d.Type.Results, m)
	r := parseField(d.Recv, m)
	if r != nil {
		funcType.Receiver = r[0]
	}
	return funcType

}

func parseGenDecl(d *ast.GenDecl, m map[string]string) *api.GoStruct {
	if d.Tok != token.TYPE {
		return nil
	}
	goStruct := api.NewGoStruct()
	for _, spec := range d.Specs {
		// goStruct := api.NewGoStruct()
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}
		if typeSpec.Name != nil && typeSpec.Name.Name != "" {
			goStruct.Name = typeSpec.Name.Name
		}
		goStruct.Field = append(goStruct.Field, parseField(structType.Fields, m)...)

	}

	return goStruct

}

func parseImport(parseFile *ast.File) map[string]string {
	importMap := make(map[string]string)
	if parseFile.Imports != nil && len(parseFile.Imports) > 0 {
		for _, spec := range parseFile.Imports {
			path := spec.Path.Value
			var name string
			if spec.Name != nil && !util.EmptyString(spec.Name.Name) {
				name = spec.Name.Name
			} else {
				pos := strings.IndexByte(path, consts.Slash)
				name = path[pos+1:]
			}
			importMap[name] = path
		}
	}
	return importMap
}

func parseField(field *ast.FieldList, importMap map[string]string) []*api.GoField {
	if field == nil && field.NumFields() == 0 {
		return nil
	}
	fields := make([]*api.GoField, 0)
	for _, f := range field.List {
		if f.Names != nil && len(f.Names) > 0 {
			for _, name := range f.Names {
				goField := api.NewGoField()
				goField.Tag = parseTag(f.Tag)
				goField.Comment = f.Comment.Text()
				goField.Open = name.IsExported()
				goField.Name = name.Name
				goField.Type = parseFieldType(f.Type, importMap)
				fields = append(fields, goField)
			}
		} else {
			goField := api.NewGoField()
			goField.Tag = parseTag(f.Tag)
			goField.Comment = f.Comment.Text()
			goField.Type = parseFieldType(f.Type, importMap)
			fields = append(fields, goField)
		}
	}
	return fields
}

func parseTag(tag *ast.BasicLit) map[string]string {
	if tag == nil || util.EmptyString(tag.Value) {
		return nil
	}
	m := make(map[string]string)

	value := tag.Value
	tagContent := util.UnwrapQuote(value)
	subtag := strings.Split(tagContent, " ")
	for _, s := range subtag {
		trimStr := strings.TrimSpace(s)
		if trimStr != "" {
			pos := strings.IndexByte(trimStr, consts.Colon)
			m[trimStr[:pos]] = util.UnwrapQuote(trimStr[pos+1:])
		}
	}
	return m
}

func parseFieldType(expr ast.Expr, importMap map[string]string) api.IFieldType {

	switch e := expr.(type) {
	case *ast.Ident:
		return api.NewIdent(e.Name)
	case *ast.SelectorExpr:
		selector := api.NewSelector()
		selector.Obj = parseFieldType(e.X, importMap)
		selector.Package = e.Sel.Name
		return selector
	case *ast.StructType:
		structType := api.NewStructType()
		structType.Field = parseField(e.Fields, importMap)
		return structType
	case *ast.Ellipsis:
		return api.NewEllipsisType(parseFieldType(e.Elt, importMap))
	case *ast.FuncType:
		funcType := api.NewFuncType()
		funcType.In = parseField(e.Params, importMap)
		funcType.Out = parseField(e.Results, importMap)
		return funcType
	case *ast.ChanType:
		chanType := api.NewChanType()
		chanType.Content = parseFieldType(e.Value, importMap)
		return chanType
	case *ast.ArrayType:
		arrayType := api.NewArrayType()
		arrayType.Content = parseFieldType(e.Elt, importMap)
		return arrayType
	case *ast.MapType:
		newMap := api.NewMap()
		newMap.Key = parseFieldType(e.Key, importMap)
		newMap.Value = parseFieldType(e.Value, importMap)
		return newMap
	case *ast.StarExpr:
		pointerType := api.NewPointerType()
		pointerType.Content = parseFieldType(e.X, importMap)
		return pointerType
	}

	return nil
}

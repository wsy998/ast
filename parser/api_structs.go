package parser

import (
	"github.com/wsy998/ast/consts"
	"github.com/wsy998/ast/util"
)

type GoStruct struct {
	Comment     string
	Package     map[string]string
	Name        string
	Field       []*GoField
	Fun         []*GoFunc
	Open        bool
	Constructor *GoFunc
}

func (g *GoStruct) String() string {
	builder := util.Text{}
	if !util.EmptyString(g.Comment) {
		builder.WriteString(consts.Comment)
		builder.WriteSpace()
		builder.WriteString(g.Comment)
	}
	builder.WriteString(consts.Type)
	builder.WriteSpace()
	builder.WriteString(g.Name)
	builder.WriteSpace()
	builder.WriteString(consts.TypeStruct)
	builder.WriteOpenBrace()
	builder.WriteEndl()
	for _, field := range g.Field {
		builder.WriteTab()
		builder.WriteStringWithEndl(field.String())
	}
	for _, field := range g.Fun {
		builder.WriteTab()
		builder.WriteStringWithEndl(field.String())
	}
	builder.WriteCloseBrace()
	builder.WriteEndl()
	return builder.String()
}

func (g *GoStruct) OpenFun() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, fun := range g.Fun {
		if fun.Open {
			funcs = append(funcs, fun)
		}
	}
	return funcs
}
func (g *GoStruct) OpenField() []*GoField {
	funcs := make([]*GoField, 0)
	for _, fun := range g.Field {
		if fun.Open {
			funcs = append(funcs, fun)
		}
	}
	return funcs
}

func NewGoStruct() *GoStruct {
	return &GoStruct{}
}

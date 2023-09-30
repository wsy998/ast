package api

import (
	"github.com/wsy998/ast/internal/consts"
	"github.com/wsy998/ast/internal/util"
)

type GoStruct struct {
	Comment     string
	Package     map[string]string
	Name        string
	Field       []*GoField
	Funs        []*GoFunc
	Open        bool
	Constructor *GoFunc
	Impl        []interface{}
}

func (g *GoStruct) String() string {

	builder := util.NewText()
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
	for _, field := range g.Funs {
		builder.WriteTab()
		builder.WriteStringWithEndl(field.String())
	}
	builder.WriteCloseBrace()
	builder.WriteEndl()
	return builder.String()
}
func NewGoStruct() *GoStruct {
	return &GoStruct{}
}

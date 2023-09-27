package parser

import (
	"github.com/wsy998/ast/internal/consts"
	"github.com/wsy998/ast/internal/util"
)

type GoField struct {
	Package string
	Name    string
	Type    string
	Field   []*GoField
	Pointer bool
	Tag     map[string]string
	Open    bool
	Comment string
}

func (f *GoField) String() string {
	builder := util.Text{}
	if !util.EmptyString(f.Name) {
		builder.WriteString(f.Name)
		builder.WriteSpace()
	}
	if f.Pointer {
		builder.WriteStar()
	}
	builder.WriteString(f.Type)

	if f.Type == consts.TypeStruct {
		builder.WriteSpace()
		builder.WriteOpenBrace()
		builder.WriteEndl()
		if f.Field != nil && len(f.Field) > 0 {
			for _, m := range f.Field {
				builder.WriteTab()
				builder.WriteTab()
				builder.WriteStringWithEndl(m.String())
			}
		}
		builder.WriteTab()
		builder.WriteCloseBrace()
		builder.WriteEndl()
	}
	if f.Tag != nil && len(f.Tag) > 0 {
		builder.WriteSpace()
		builder.WriteTagSign()
		first := true
		for s, s2 := range f.Tag {
			if !first {
				builder.WriteSpace()
			} else {
				first = true
			}
			builder.WriteString(s)
			builder.WriteByte(consts.Colon)
			builder.WriteWithQuote(s2)
		}
		builder.WriteTagSign()
	}
	return builder.String()
}

func NewGoField() *GoField {
	return &GoField{}
}

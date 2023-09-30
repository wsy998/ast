package api

import (
	"github.com/wsy998/ast/internal/consts"
	"github.com/wsy998/ast/internal/util"
)

type GoField struct {
	Open    bool
	Tag     map[string]string
	Name    string
	Type    IFieldType
	Comment string
}

func (f *GoField) String() string {
	builder := util.NewText()
	if !util.EmptyString(f.Name) {
		builder.WriteString(f.Name)
		builder.WriteSpace()
	}
	if f.Type != nil {
		builder.WriteString(f.Type.String())
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

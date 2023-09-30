package api

import (
	"github.com/wsy998/ast/consts"
	"github.com/wsy998/ast/internal/util"
)

type FuncType struct {
	Name string
	In   []*GoField
	Out  []*GoField
}

func (f *FuncType) String() string {
	text := util.NewText()
	text.WriteString("func ")

	if !util.EmptyString(f.Name) {
		text.WriteString(f.Name)
	}
	text.WriteOpenParen()
	if f.In != nil && len(f.In) > 0 {
		for _, in := range f.In {
			text.WriteString(in.String())
		}
	}
	text.WriteCloseParen()
	if f.Out != nil && len(f.Out) > 0 {
		for _, out := range f.Out {
			text.WriteString(out.String())
		}
	}
	text.WriteOpenBrace()
	text.WriteCloseBrace()
	return text.String()
}

func (f *FuncType) Pointer() bool {
	return false
}

func (f *FuncType) Chan() bool {
	return false
}

func (f *FuncType) TypeName() string {
	return consts.Func
}

func (f *FuncType) RealType() IFieldType {
	return nil
}

func NewFuncType() *FuncType {
	return &FuncType{}
}

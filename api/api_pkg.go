package api

import (
	"strings"

	"github.com/wsy998/ast/internal/consts"
	"github.com/wsy998/ast/internal/util"
)

type GoPkg struct {
	Name      string
	GoStructs []*GoStruct
	Func      []*GoFunc
	Imports   map[string]string
}

func (f *GoPkg) String() string {

	text := util.NewText()
	text.WriteString("package " + f.Name)
	text.WriteEndl()
	if f.Imports != nil && len(f.Imports) > 0 {
		for n, p := range f.Imports {
			if strings.HasSuffix(p, n) {
				text.WriteString(p)
			} else {
				text.WriteString(n + " " + p)
			}
		}
	}

	for _, goStruct := range f.GoStructs {
		text.WriteString(consts.Type)
		text.WriteSpace()
		text.WriteString(goStruct.Name)
		text.WriteOpenBrace()
		if goStruct.Field != nil {
			text.WriteEndl()
			for _, field := range goStruct.Field {
				field.String()
			}
			text.WriteCloseBrace()
		}
		text.WriteEndl()
	}
	for _, goFunc := range f.Func {
		text.WriteStringWithEndl(goFunc.String())
	}
	return text.String()
}

func NewGoPkg() *GoPkg {
	return &GoPkg{}
}

func (f *GoPkg) WithoutReceiver() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, goFunc := range f.Func {
		if goFunc.Receiver == nil {
			funcs = append(funcs, goFunc)
		}
	}
	return funcs
}

func (f *GoPkg) OpenReceiver() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, goFunc := range f.Func {
		if goFunc.Open {
			funcs = append(funcs, goFunc)
		}
	}
	return funcs
}
func (f *GoPkg) OpenWithoutReceiver() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, goFunc := range f.Func {
		if goFunc.Open && goFunc.Receiver != nil {
			funcs = append(funcs, goFunc)
		}
	}
	return funcs
}

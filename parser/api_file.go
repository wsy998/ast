package parser

import (
	"strings"

	"github.com/wsy998/ast/consts"
	"github.com/wsy998/ast/util"
)

type GoFile struct {
	GoStructs []*GoStruct
	Func      []*GoFunc
	Imports   map[string]string
}

func (f *GoFile) String() string {
	builder := util.Text{}
	if len(f.Imports) > 0 {
		builder.WriteString(consts.Import)
		if len(f.Imports) > 1 {
			builder.WriteSpace()
			builder.WriteOpenParen()
			for n, v := range f.Imports {
				if !util.EmptyString(n) {
					builder.WriteSpace()
					builder.WriteString(n)
				}
				builder.WriteSpace()
				builder.WriteWithQuote(v)
				builder.WriteEndl()
			}
			builder.WriteStringWithEndl(string(consts.CloseParen))
		} else {
			for n, v := range f.Imports {
				if !strings.HasSuffix(v, n) {
					builder.WriteSpace()
					builder.WriteString(n)
				}
				builder.WriteSpace()
				builder.WriteWithQuote(v)
				builder.WriteEndl()
			}
		}
	}
	for _, structItem := range f.GoStructs {
		builder.WriteStringWithEndl(structItem.String())
	}
	for _, funcItem := range f.Func {
		builder.WriteStringWithEndl(funcItem.String())
	}
	return builder.String()
}

func NewGoFile() *GoFile {
	return &GoFile{}
}

func (f *GoFile) WithoutReceiver() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, goFunc := range f.Func {
		if goFunc.Receiver == nil {
			funcs = append(funcs, goFunc)
		}
	}
	return funcs
}

func (f *GoFile) OpenReceiver() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, goFunc := range f.Func {
		if goFunc.Open {
			funcs = append(funcs, goFunc)
		}
	}
	return funcs
}
func (f *GoFile) OpenWithoutReceiver() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, goFunc := range f.Func {
		if goFunc.Open && goFunc.Receiver != nil {
			funcs = append(funcs, goFunc)
		}
	}
	return funcs
}

package parser

import (
	"strings"

	"github.com/wsy998/ast/internal/consts"
	"github.com/wsy998/ast/internal/util"
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

// WithoutReceiver Get all functions without receivers.
func (f *GoFile) WithoutReceiver() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, goFunc := range f.Func {
		if goFunc.Receiver == nil {
			funcs = append(funcs, goFunc)
		}
	}
	return funcs
}

// ExportedFunc Get all exported functions.
func (f *GoFile) ExportedFunc() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, goFunc := range f.Func {
		if goFunc.Open {
			funcs = append(funcs, goFunc)
		}
	}
	return funcs
}

// ExportedWithoutReceiver Get exported functions without receivers.
func (f *GoFile) ExportedWithoutReceiver() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, goFunc := range f.Func {
		if goFunc.Open && goFunc.Receiver != nil {
			funcs = append(funcs, goFunc)
		}
	}
	return funcs
}

// FuncByName Retrieve functions by name, regardless of whether they have a receiver.
func (f *GoFile) FuncByName(name string) *GoFunc {
	for _, goFunc := range f.Func {
		if goFunc.Name == name {
			return goFunc
		}
	}
	return nil
}

// FuncByNameWithoutReceiver Retrieve functions by name, but they must not have a receiver.
func (f *GoFile) FuncByNameWithoutReceiver(name string) *GoFunc {
	for _, goFunc := range f.Func {
		if goFunc.Name == name && goFunc.Receiver == nil {
			return goFunc
		}
	}
	return nil
}

// FuncByNameWithReceiver Retrieve functions by name, but they must have a receiver.
func (f *GoFile) FuncByNameWithReceiver(name string) *GoFunc {
	for _, goFunc := range f.Func {
		if goFunc.Name == name && goFunc.Receiver != nil {
			return goFunc
		}
	}
	return nil
}

// ExportedFuncByName Retrieve exported functions by name.
func (f *GoFile) ExportedFuncByName(name string) *GoFunc {
	for _, goFunc := range f.Func {
		if goFunc.Name == name && goFunc.Open {
			return goFunc
		}
	}
	return nil
}

// ExportedFuncByNameWithoutReceiver Retrieve exported functions by name, but they cannot have a receiver.
func (f *GoFile) ExportedFuncByNameWithoutReceiver(name string) *GoFunc {
	for _, goFunc := range f.Func {
		if goFunc.Name == name && goFunc.Receiver == nil {
			return goFunc
		}
	}
	return nil
}

// ExportedFuncByNameWithReceiver Retrieve exported functions by name, but they must have a receiver.
func (f *GoFile) ExportedFuncByNameWithReceiver(name string) *GoFunc {
	for _, goFunc := range f.Func {
		if goFunc.Name == name && goFunc.Receiver != nil {
			return goFunc
		}
	}
	return nil
}

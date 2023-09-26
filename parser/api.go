package parser

import (
	"fmt"
	"strings"
)

type GoField struct {
	Package map[string]string
	Name    string
	Type    string
	Field   []*GoField
	Pointer bool
	Tag     map[string]string
	Open    bool
}

func (f *GoField) String() string {
	builder := strings.Builder{}
	if f.Name != "" {
		builder.WriteString(f.Name + " ")
	}
	if f.Pointer {
		builder.WriteByte('*')
	}
	builder.WriteString(f.Type)
	if f.Tag != nil && len(f.Tag) > 0 {
		builder.WriteString("`")
		for s, s2 := range f.Tag {
			builder.WriteString(fmt.Sprintf("%s:\"%s\" ", s, s2))
		}
		builder.WriteString("`")

	}
	return builder.String()
}

func NewGoField() *GoField {
	return &GoField{}
}

type GoFunc struct {
	Package  map[string]string
	Receiver *GoField
	Comment  string
	Name     string
	In       []*GoField
	Out      []*GoField
	Open     bool
}

func (funcS *GoFunc) String() string {
	builder := strings.Builder{}
	if funcS.Comment != "" {
		builder.WriteString(fmt.Sprintf("// %s\n", funcS.Comment))
	}
	builder.WriteString("func ")
	if funcS.Receiver != nil {
		builder.WriteString(fmt.Sprintf("(%s)", funcS.Receiver.String()))
	}
	builder.WriteString(funcS.Name)
	builder.WriteByte('(')
	if len(funcS.In) > 0 {
		for i, field := range funcS.In {
			if i != 0 {
				builder.WriteByte(',')
			}
			builder.WriteString(field.String())
		}
	}
	builder.WriteByte(')')
	if len(funcS.Out) > 0 {
		if funcS.Out[0].Name != "" {
			builder.WriteByte('(')
			for i, o := range funcS.Out {
				if i != 0 {
					builder.WriteByte(',')
				}
				builder.WriteString(o.String())
			}
			builder.WriteByte(')')
		} else if len(funcS.Out) > 1 {
			builder.WriteByte('(')
			for i, o := range funcS.Out {
				if i != 0 {
					builder.WriteByte(',')
				}
				builder.WriteString(o.String())
			}
			builder.WriteByte(')')
		} else {
			builder.WriteString(funcS.Out[0].Type)
		}
	}
	return builder.String()
}

func NewGoFunc() *GoFunc {
	return &GoFunc{}
}

type GoStruct struct {
	Comment string
	Package map[string]string
	Name    string
	Field   []*GoField
	Fun     []*GoFunc
	Open    bool
}

func (g *GoStruct) String() string {
	builder := strings.Builder{}
	if g.Comment != "" {
		builder.WriteString(fmt.Sprintf("// %s", g.Comment))
	}
	builder.WriteString(fmt.Sprintf("type %s struct {\n", g.Name))
	for _, field := range g.Field {
		builder.WriteString("\t" + field.String() + "\n")
	}
	for _, field := range g.Fun {
		builder.WriteString("\t" + field.String() + "\n")
	}
	builder.WriteString("}\n")
	return builder.String()
}

func NewGoStruct() *GoStruct {
	return &GoStruct{}
}

type GoFile struct {
	Package   map[string]string
	GoStructs []*GoStruct
	Func      []*GoFunc
}

func (g *GoFile) String() string {
	builder := strings.Builder{}
	for _, f := range g.GoStructs {
		builder.WriteString(f.String() + "\n")
	}
	for _, f := range g.Func {
		builder.WriteString(f.String() + "\n")
	}
	return builder.String()
}

func NewGoFile() *GoFile {
	return &GoFile{}
}

type GoPkg struct {
	FilePkg   []*GoFile
	GoStructs []*GoStruct
	Func      []*GoFunc
}

func NewGoPkg() *GoPkg {
	return &GoPkg{}
}

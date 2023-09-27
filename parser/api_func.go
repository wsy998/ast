package parser

import (
	"github.com/wsy998/ast/consts"
	"github.com/wsy998/ast/util"
)

type GoFunc struct {
	Package  map[string]string
	Receiver *GoField
	Comment  string
	Name     string
	In       []*GoField
	Out      []*GoField
	Open     bool
}

func (g *GoFunc) String() string {
	builder := util.Text{}
	if !util.EmptyString(g.Comment) {
		builder.WriteString(consts.Comment)
		builder.WriteSpace()
		builder.WriteString(g.Comment)
		builder.WriteEndl()
	}
	builder.WriteString(consts.Func)
	builder.WriteSpace()
	if g.Receiver != nil {
		builder.Writef("(%s)", g.Receiver.String())
	}
	builder.WriteString(g.Name)
	builder.WriteOpenParen()
	if len(g.In) > 0 {
		for i, field := range g.In {
			if i != 0 {
				builder.WriteComma()
			}
			builder.WriteString(field.String())
		}
	}
	builder.WriteCloseParen()
	if len(g.Out) > 0 {
		if !util.EmptyString(g.Out[0].Name) {
			builder.WriteOpenParen()
			for i, o := range g.Out {
				if i != 0 {
					builder.WriteComma()
				}
				builder.WriteString(o.String())
			}
			builder.WriteCloseParen()
		} else if len(g.Out) > 1 {
			builder.WriteOpenParen()
			for i, o := range g.Out {
				if i != 0 {
					builder.WriteComma()
				}
				builder.WriteString(o.String())
			}
			builder.WriteCloseParen()
		} else {
			builder.WriteString(g.Out[0].Type)
		}
	}
	return builder.String()
}

func NewGoFunc() *GoFunc {
	return &GoFunc{}
}

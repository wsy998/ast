package api

import (
	"github.com/wsy998/ast/internal/consts"
	"github.com/wsy998/ast/internal/util"
)

type GoFunc struct {
	Open     bool
	In       []*GoField
	Out      []*GoField
	Package  map[string]string
	Receiver *GoField
	Comment  string
	Name     string
}

func (g *GoFunc) String() string {

	builder := util.NewText()
	if !util.EmptyString(g.Comment) {
		builder.WriteString(consts.Comment)
		builder.WriteSpace()
		builder.WriteString(g.Comment)
		builder.WriteEndl()
	}
	builder.WriteString(consts.Func)
	if g.Receiver != nil {
		builder.WriteSpace()
		builder.Writef("(%s)", g.Receiver.String())
	}
	builder.WriteSpace()
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
			builder.WriteString(g.Out[0].String())
		}
	}
	return builder.String()
}

func NewGoFunc() *GoFunc {
	return &GoFunc{}
}

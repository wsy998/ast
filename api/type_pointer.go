package api

import (
	"github.com/wsy998/ast/consts"
)

type PointerType struct {
	Content IFieldType
}

func NewPointerType() *PointerType {
	return &PointerType{}
}

func (p *PointerType) String() string {
	return "*" + p.RealType().String()
}

func (p *PointerType) Pointer() bool {
	return true
}

func (p *PointerType) Chan() bool {
	return false
}

func (p *PointerType) TypeName() string {
	return consts.Pointer
}

func (p *PointerType) RealType() IFieldType {
	return p.Content
}

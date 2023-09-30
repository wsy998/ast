package api

import (
	"github.com/wsy998/ast/consts"
)

type ChanType struct {
	Content IFieldType
}

func (c *ChanType) String() string {
	return "chan " + c.Content.String()
}

func (c *ChanType) Pointer() bool {
	return false
}

func (c *ChanType) Chan() bool {
	return true
}

func (c *ChanType) TypeName() string {
	return consts.Chan
}

func (c *ChanType) RealType() IFieldType {
	return c.Content
}

func NewChanType() *ChanType {
	return &ChanType{}
}

package api

import (
	api_const "github.com/wsy998/ast/consts"
	"github.com/wsy998/ast/internal/consts"
	"github.com/wsy998/ast/internal/util"
)

type StructType struct {
	Name  string
	Field []*GoField
}

func (s *StructType) String() string {
	text := util.NewText()
	text.WriteString(consts.TypeStruct)
	text.WriteSpace()

	if s.Name != "" {
		text.WriteString(s.Name)
		text.WriteSpace()
	}
	text.WriteOpenBrace()
	if s.Field != nil && len(s.Field) > 0 {
		text.WriteEndl()
		for _, fieldType := range s.Field {
			text.WriteTab()
			text.WriteTab()
			text.WriteString(fieldType.Name)
			text.WriteSpace()
			text.WriteString(fieldType.Type.String())
			text.WriteEndl()
		}
		text.WriteTab()
	}
	text.WriteCloseBrace()

	return text.String()
}

func (s *StructType) Pointer() bool {
	return false
}

func (s *StructType) Chan() bool {
	return false
}

func (s *StructType) TypeName() string {
	return api_const.Struct
}

func (s *StructType) RealType() IFieldType {
	return nil
}

func NewStructType() *StructType {
	return &StructType{}
}

package api

import (
	"fmt"
)

type Map struct {
	Key   IFieldType
	Value IFieldType
}

func (m Map) String() string {
	return fmt.Sprintf("map[%s]%s", m.Key.String(), m.Value.String())
}

func (m Map) Pointer() bool {
	return false

}

func (m Map) Chan() bool {
	return false
}

func (m Map) TypeName() string {
	return "Map"
}

func (m Map) RealType() IFieldType {
	return nil
}

func NewMap() *Map {
	return &Map{}
}

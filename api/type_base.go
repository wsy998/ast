package api

import (
	"fmt"
)

type IFieldType interface {
	fmt.Stringer
	Pointer() bool
	Chan() bool
	TypeName() string
	RealType() IFieldType
}

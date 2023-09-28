package parser

import (
	"reflect"
)

func init() {
	UnsafePointer
	reflect.ValueOf().Kind()
}

const (
	Bool          = "bool"
	Int           = "int"
	Int8          = "int8"
	Int16         = "int16"
	Int32         = "int32"
	Int64         = "int64"
	Uint          = "uint"
	Uint8         = "uint8"
	Uint16        = "uint16"
	Uint32        = "uint32"
	Uint64        = "uint64"
	Uintptr       = "uintptr"
	Float32       = "float32"
	Float64       = "float64"
	Complex64     = "complex64"
	Complex128    = "complex128"
	Array         = "array"
	Chan          = "chan"
	Func          = "func"
	Interface     = "uint8"
	Map           = "map"
	Pointer       = "pointer"
	Slice         = "slice"
	String        = "string"
	Struct        = "struct"
	UnsafePointer = "unsafePointer"
)

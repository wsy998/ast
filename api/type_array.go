package api

type ArrayType struct {
	Content IFieldType
}

func (a *ArrayType) String() string {
	return "[]" + a.Content.String()
}

func (a *ArrayType) Pointer() bool {
	return false
}

func (a *ArrayType) Chan() bool {
	return false
}

func (a *ArrayType) TypeName() string {
	return "array"
}

func (a *ArrayType) RealType() IFieldType {
	return a.Content
}

func NewArrayType() *ArrayType {
	return &ArrayType{}
}

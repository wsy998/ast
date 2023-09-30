package api

type EllipsisType struct {
	Content IFieldType
}

func NewEllipsisType(content IFieldType) *EllipsisType {
	return &EllipsisType{Content: content}
}

func (e *EllipsisType) String() string {
	return "..." + e.Content.String()
}

func (e *EllipsisType) Pointer() bool {
	return false
}

func (e *EllipsisType) Chan() bool {
	return false
}

func (e *EllipsisType) TypeName() string {
	return "ellipsis"
}

func (e *EllipsisType) RealType() IFieldType {
	return e.Content

}

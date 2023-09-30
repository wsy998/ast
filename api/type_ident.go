package api

type IdentType struct {
	Name string
}

func (i *IdentType) String() string {
	return i.Name
}

func (i *IdentType) Pointer() bool {
	return false
}

func (i *IdentType) Chan() bool {
	return false
}

func (i *IdentType) TypeName() string {
	return "ident"
}

func (i *IdentType) RealType() IFieldType {
	return nil
}

func NewIdent(name string) *IdentType {
	return &IdentType{Name: name}
}

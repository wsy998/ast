package api

type SelectorType struct {
	Package string
	Obj     IFieldType
}

func (s *SelectorType) String() string {
	return s.Package + "." + s.Obj.String()
}

func (s *SelectorType) Pointer() bool {
	return false
}

func (s *SelectorType) Chan() bool {
	return false
}

func (s *SelectorType) TypeName() string {
	return "Selector"
}

func (s *SelectorType) RealType() IFieldType {
	return nil
}

func NewSelector() *SelectorType {
	return &SelectorType{}
}

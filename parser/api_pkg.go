package parser

type GoPkg struct {
	GoStructs []*GoStruct
	Func      []*GoFunc
	Imports   map[string]string
}

func NewGoPkg() *GoPkg {
	return &GoPkg{}
}

func (f *GoPkg) WithoutReceiver() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, goFunc := range f.Func {
		if goFunc.Receiver == nil {
			funcs = append(funcs, goFunc)
		}
	}
	return funcs
}

func (f *GoPkg) OpenReceiver() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, goFunc := range f.Func {
		if goFunc.Open {
			funcs = append(funcs, goFunc)
		}
	}
	return funcs
}
func (f *GoPkg) OpenWithoutReceiver() []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, goFunc := range f.Func {
		if goFunc.Open && goFunc.Receiver != nil {
			funcs = append(funcs, goFunc)
		}
	}
	return funcs
}

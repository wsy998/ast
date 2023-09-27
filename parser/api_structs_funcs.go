package parser

// AllExportedFun Retrieve all exported functions.
func (g *GoStruct) AllExportedFun() []*GoFunc {
	return g.GetFuncsByFilter(func(f *GoFunc) bool {
		return f.Open
	})
}

// AllNotExportedFun Retrieve all non-exported functions.
func (g *GoStruct) AllNotExportedFun() []*GoFunc {
	return g.GetFuncsByFilter(func(f *GoFunc) bool {
		return !f.Open
	})
}

// GetFunByName Retrieve functions by name, regardless of whether they are exported.
func (g *GoStruct) GetFunByName(name string) *GoFunc {
	return g.GetFuncByFilter(func(f *GoFunc) bool {
		return f.Name == name
	})
}

// GetExportedFunByName Retrieve functions by name, but the functions must be exported.
func (g *GoStruct) GetExportedFunByName(name string) *GoFunc {
	return g.GetFuncByFilter(func(f *GoFunc) bool {
		return f.Name == name && f.Open
	})
}

// GetFunByNameNotExport Retrieve functions by name, but the functions must not be exported.
func (g *GoStruct) GetFunByNameNotExport(name string) *GoFunc {
	return g.GetFuncByFilter(func(f *GoFunc) bool {
		return f.Name == name && !f.Open
	})
}

// GetFuncsByFilter Retrieve multiple functions based on another function.
func (g *GoStruct) GetFuncsByFilter(s func(f *GoFunc) bool) []*GoFunc {
	funcs := make([]*GoFunc, 0)
	for _, fun := range g.Funs {
		if s(fun) {
			funcs = append(funcs, fun)
		}
	}
	return funcs
}

// GetFuncByFilter Retrieve a single function based on another function.
func (g *GoStruct) GetFuncByFilter(s func(f *GoFunc) bool) *GoFunc {
	for _, fun := range g.Funs {
		if s(fun) {
			return fun
		}
	}
	return nil
}

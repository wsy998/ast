package parser

// AllExportedField Retrieve all exported fields.
func (g *GoStruct) AllExportedField() []*GoField {
	return g.GetFieldsByFilter(func(f *GoField) bool {
		return f.Open
	})
}

// AllNotExportedField Retrieve all non-exported fields.
func (g *GoStruct) AllNotExportedField() []*GoField {
	return g.GetFieldsByFilter(func(f *GoField) bool {
		return !f.Open
	})
}

// GetFieldByName Retrieve field by name, regardless of whether they are exported.
func (g *GoStruct) GetFieldByName(name string) *GoField {
	return g.GetFieldByFilter(func(f *GoField) bool {
		return f.Name == name
	})
}

// GetExportedFieldByName Retrieve field by name, but the field must be exported.
func (g *GoStruct) GetExportedFieldByName(name string) *GoField {
	return g.GetFieldByFilter(func(f *GoField) bool {
		return f.Name == name && f.Open
	})
}

// GetFieldByNameNotExport Retrieve field by name, but the field must not be exported.
func (g *GoStruct) GetFieldByNameNotExport(name string) *GoField {
	return g.GetFieldByFilter(func(f *GoField) bool {
		return f.Name == name && !f.Open
	})
}

// GetFieldsByFilter Retrieve multiple fields based on another function."
func (g *GoStruct) GetFieldsByFilter(s func(f *GoField) bool) []*GoField {
	funcs := make([]*GoField, 0)
	for _, fun := range g.Field {
		if s(fun) {
			funcs = append(funcs, fun)
		}
	}
	return funcs
}

// GetFieldByFilter Retrieve a single field based on a field name.
func (g *GoStruct) GetFieldByFilter(s func(f *GoField) bool) *GoField {
	for _, fun := range g.Field {
		if s(fun) {
			return fun
		}
	}
	return nil
}

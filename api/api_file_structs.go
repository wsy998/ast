package api

// GetStructByName Retrieve a struct by function, regardless of whether it's exported.
func (f *GoFile) GetStructByName(name string) *GoStruct {
	return f.GetStructByFilter(func(goStruct *GoStruct) bool {
		return goStruct.Name == name
	})
}

// GetAllExportedFunc Retrieve all exported structs.
func (f *GoFile) GetAllExportedFunc() []*GoStruct {
	return f.GetStructsByFilter(func(goStruct *GoStruct) bool {
		return goStruct.Open
	})
}

// GetStructByNameWithoutExport Retrieve non-exported structs.
func (f *GoFile) GetStructByNameWithoutExport(name string) *GoStruct {
	return f.GetStructByFilter(func(goStruct *GoStruct) bool {
		return !goStruct.Open && goStruct.Name == name
	})
}

// GetStructByNameWithExport Retrieve exported structs.
func (f *GoFile) GetStructByNameWithExport(name string) *GoStruct {
	return f.GetStructByFilter(func(goStruct *GoStruct) bool {
		return goStruct.Open && goStruct.Name == name
	})
}

// GetStructByFilter Retrieve a single struct based on another function.
func (f *GoFile) GetStructByFilter(filter func(goStruct *GoStruct) bool) *GoStruct {
	for _, goStruct := range f.GoStructs {
		if filter(goStruct) {
			return goStruct
		}
	}
	return nil
}

// GetStructsByFilter Retrieve multiple structs based on another function.
func (f *GoFile) GetStructsByFilter(filter func(goStruct *GoStruct) bool) []*GoStruct {
	structs := make([]*GoStruct, 0)
	for _, goStruct := range f.GoStructs {
		if filter(goStruct) {
			structs = append(structs, goStruct)
		}
	}
	return structs
}

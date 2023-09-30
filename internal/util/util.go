package util

import (
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

func ParseGoMod(path string) (string, error) {
	goModPath := filepath.Join(path, "go.mod")
	file, err := os.ReadFile(goModPath)
	if err != nil {
		return "", err
	}
	parse, err := modfile.Parse(goModPath, file, nil)
	if err != nil {
		return "", err
	}
	return parse.Module.Mod.Path, nil
}

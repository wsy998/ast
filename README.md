# ast

Based on the official AST package, support single files and single packages, but currently do not support package-level
variables and struct composition.

## Installation

`go get -u github.com/wsy998/ast`

## Usage
```go
func Parse(file string) (*GoFile, error){} // parse single file
func ParsePackage(pkg string) (*GoPkg, error){} // parse single package

```

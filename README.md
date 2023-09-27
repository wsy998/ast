# ast

Based on the official AST package, support single files and single packages, but currently do not support package-level
variables and struct composition.
[中文文档](./README_ZH_CN.md)
## Features
Parse a single file or single package, and retrieve all the structs and functions within the file or package.


## Installation

`go get -u github.com/wsy998/ast`

## Usage
```go
func Parse(file string) (*GoFile, error){} // parse single file
func ParsePackage(pkg string) (*GoPkg, error){} // parse single package

```

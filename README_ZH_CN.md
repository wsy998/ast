# ast

基于官方的AST封装，支持单个文件，单个包，但暂时不支持包级别的变量和结构体组合。

## 功能

解析单个文件或单个包，获取文件或包内的所有结构体以及所有的函数

## 安装

`go get -u github.com/wsy998/ast`

## 用法

```go
func Parse(file string) (*GoFile, error){} // parse single file
func ParsePackage(pkg string) (*GoPkg, error){} // parse single package

```

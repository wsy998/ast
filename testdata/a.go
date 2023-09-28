package testdata

import (
	"os"
)

func A() {

}

func B() {
}

func C(a string) {
}
func D(a string) (b string) {
	return ""
}
func E(a *string) {
}
func F(a *string) (c *string) {
	l := "123"
	return &l
}
func G() (c *os.File) {
	return &os.File{}
}
func H() (c func()) {
	var s = func() {}
	return s
}
func I(c os.File) {
}
func J(c *os.File) {
}

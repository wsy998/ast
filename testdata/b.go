package testdata

import (
	"os"
)

type l struct {
	Q string
}

type O struct {
	l
	a string `k:"ss"`
	c chan *string
	g chan string
	d os.File
	e *os.File
	f struct {
		A string
		b int
	}
}

func (receiver *O) A() {

}

func (receiver *O) B() {
}

func (receiver *O) C(a string) {
}
func (receiver *O) D(a string) (b string) {
	return ""
}
func (receiver *O) E(a *string) {
}
func (receiver *O) F(a *string) (c *string) {
	l := "123"
	return &l
}
func (receiver *O) G() (c *os.File) {
	return &os.File{}
}
func (receiver *O) H() (c os.File) {
	return os.File{}
}
func (receiver *O) I(c os.File) {
}
func (receiver *O) J(c *os.File) {
}

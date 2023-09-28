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
func makeVertexShader(src string) uint32 {
	return 0
	// vShader := gl.CreateShader(gl.VERTEX_SHADER)
	//
	// // str := gl.Str(src)
	// strs, free := gl.Strs(src)
	// defer free()
	// // _ = free
	// // free()
	// gl.ShaderSource(vShader, 1, strs, nil)
	// gl.CompileShader(vShader)
	// var status int32
	// gl.GetShaderiv(vShader, gl.COMPILE_STATUS, &status)
	// log := [512]byte{}
	// if status == gl.FALSE {
	// 	logLength := int32(0)
	// 	gl.GetShaderInfoLog(vShader, 512, &logLength, &log[0])
	// 	fmt.Printf("VertexShader failed to compile %v: %s\n", src, string(log[:logLength]))
	// }

	// return vShader
}

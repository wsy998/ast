package testdata

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func makeVertexShader(src string) uint32 {
	vShader := gl.CreateShader(gl.VERTEX_SHADER)

	// str := gl.Str(src)
	strs, free := gl.Strs(src)
	defer free()
	// _ = free
	// free()
	gl.ShaderSource(vShader, 1, strs, nil)
	gl.CompileShader(vShader)
	var status int32
	gl.GetShaderiv(vShader, gl.COMPILE_STATUS, &status)
	log := [512]byte{}
	if status == gl.FALSE {
		logLength := int32(0)
		gl.GetShaderInfoLog(vShader, 512, &logLength, &log[0])
		fmt.Printf("VertexShader failed to compile %v: %s\n", src, string(log[:logLength]))
	}

	return vShader
}

func makeFragmentShader(src string) uint32 {
	fShader := gl.CreateShader(gl.FRAGMENT_SHADER)

	// str := gl.Str(src)

	strs, free := gl.Strs(src)
	defer free()
	gl.ShaderSource(fShader, 1, strs, nil)
	gl.CompileShader(fShader)

	var status int32
	gl.GetShaderiv(fShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		log := [512]byte{}
		gl.GetShaderInfoLog(fShader, 512, &logLength, &log[0])
		fmt.Printf("FragmentShader failed to compile %v: %s\n", src, string(log[:logLength]))
	}
	return fShader
}

func makeShaderProgram(shaders ...uint32) uint32 {

	shaderProg := gl.CreateProgram()

	for _, shader := range shaders {
		gl.AttachShader(shaderProg, shader)

	}
	gl.LinkProgram(shaderProg)

	var status int32
	gl.GetProgramiv(shaderProg, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		log := [512]byte{}

		gl.GetProgramInfoLog(shaderProg, 512, &logLength, &log[0])
		fmt.Printf("failed to link: %s\n", string(log[:logLength]))
	}

	return shaderProg
}

type Shader struct {
	// 程序ID
	id uint32
}

// 构造器读取并构建着色器
func NewShader(vertexPath, fragmentPath string) *Shader {
	vf, err := os.Open(vertexPath)
	if err != nil {
		panic(fmt.Errorf("打开vertex shader文件错误 %v", err))
	}
	defer vf.Close()
	ff, err := os.Open(fragmentPath)
	if err != nil {
		panic(fmt.Errorf("打开fragment shader文件错误 %v", err))
	}
	defer ff.Close()

	vfs := strings.Builder{}
	ffs := strings.Builder{}

	vbytes, err := io.ReadAll(vf)
	if err != nil {
		panic(fmt.Errorf("读取vertex shader文件错误 %v", err))
	}
	fbytes, err := io.ReadAll(ff)
	if err != nil {
		panic(fmt.Errorf("读取fragment shader文件错误 %v", err))
	}

	vfs.Write(vbytes)
	vfs.WriteString("\x00")
	ffs.Write(fbytes)
	ffs.WriteString("\x00")
	vShader := makeVertexShader(vfs.String())
	fShader := makeFragmentShader(ffs.String())
	defer gl.DeleteShader(vShader)
	defer gl.DeleteShader(fShader)

	prog := makeShaderProgram(vShader, fShader)
	return &Shader{
		id: prog,
	}
}

// 使用/激活程序
func (s *Shader) Use() {
	gl.UseProgram(s.id)
}

// uniform工具函数
func (s *Shader) SetBool(name string, value bool) {
	location := gl.GetUniformLocation(s.id, gl.Str(Str2Cstr(name)))
	v := int32(0)
	if value {
		v = 1
	}
	gl.Uniform1i(location, v)
}
func (s *Shader) SetInt(name string, value int) {
	location := gl.GetUniformLocation(s.id, gl.Str(Str2Cstr(name)))

	gl.Uniform1i(location, int32(value))
}
func (s *Shader) SetFloat(name string, value float32) {
	location := gl.GetUniformLocation(s.id, gl.Str(Str2Cstr(name)))

	gl.Uniform1f(location, value)
}

func (s *Shader) SetMat4(name string, value *float32) {
	location := gl.GetUniformLocation(s.id, gl.Str(Str2Cstr(name)))

	gl.UniformMatrix4fv(location, 1, false, value)
}

func (s *Shader) Free() {
	gl.DeleteProgram(s.id)
}

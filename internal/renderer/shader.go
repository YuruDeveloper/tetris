package renderer

import (
	packagederror "gitea.bytedev.duckdns.org/tetris/internal/packagedError"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var VertexShaderSource , _ = gl.Strs("")
var FragmentShaderSource , _ = gl.Strs("")

type Shaders struct{
	program uint32
}

func (instance *Shaders) CompileShaders() error {
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vertexShader,1,VertexShaderSource,nil);
	gl.CompileShader(vertexShader)
	err := instance.checkShaderStatus(vertexShader)
	if err != nil {
		return err
	}
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(fragmentShader,1,FragmentShaderSource,nil)
	gl.CompileShader(fragmentShader)
	err = instance.checkShaderStatus(fragmentShader)
	if err != nil {
		return err
	}
	instance.program = gl.CreateProgram()
	gl.AttachShader(instance.program,vertexShader)
	gl.AttachShader(instance.program,fragmentShader)
	gl.LinkProgram(instance.program)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	return instance.checkProgramStatus(instance.program)
}

func (instance *Shaders) checkShaderStatus(Shader uint32) error {
	var isCompile int32
	gl.GetShaderiv(Shader,gl.COMPILE_STATUS,&isCompile)
	if isCompile == gl.FALSE {
		var maxLength int32
		gl.GetShaderiv(Shader,gl.INFO_LOG_LENGTH,&maxLength)
		var information *uint8
		gl.GetShaderInfoLog(Shader,maxLength,&maxLength,information)
		err := gl.GoStr(information)
		return packagederror.NewError(packagederror.ShaderCompileFail,err)
	}	
	return nil
}


func (instance *Shaders) checkProgramStatus(program uint32) error {
	var isLinked int32
	gl.GetProgramiv(program,gl.LINK_STATUS,&isLinked)
	if isLinked == gl.FALSE {
		var maxLength int32
		gl.GetProgramiv(program,gl.INFO_LOG_LENGTH,&maxLength)
		var information *uint8
		gl.GetProgramInfoLog(program,maxLength,&maxLength,information)
		err := gl.GoStr(information)
		return packagederror.NewError(packagederror.ProgramLinkFail,err)
	}	
	return nil
}

func (instance *Shaders) GetProgram() uint32 {
	return instance.program
}
package renderer

import (
	packagederror "gitea.bytedev.duckdns.org/tetris/internal/packagedError"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var VertexShaderSource , FreeVertexShaderSource = gl.Strs("")
var FragmentShaderSource , FreeFragmentShaderSource = gl.Strs("")

type Shaders struct{
	program uint32
}

func (instance *Shaders) CompileShaders() error {
	if gl.IsProgram(instance.program) {
		return packagederror.NewError(packagederror.AlreadyCompiled,"Already Compiled") 
	}
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
	err = instance.checkProgramStatus(instance.program)
	if err == nil {
		FreeVertexShaderSource()
		FreeFragmentShaderSource()
	}
	return err
}

func (instance *Shaders) checkShaderStatus(Shader uint32) error {
	var isCompile int32
	gl.GetShaderiv(Shader,gl.COMPILE_STATUS,&isCompile)
	if isCompile == gl.FALSE {
		var maxLength int32
		gl.GetShaderiv(Shader,gl.INFO_LOG_LENGTH,&maxLength)
		information :=  make([]uint8,maxLength)
		gl.GetShaderInfoLog(Shader,maxLength,&maxLength,&information[0])
		err := gl.GoStr(&information[0])
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
		information :=  make([]uint8,maxLength)
		gl.GetProgramInfoLog(program,maxLength,&maxLength,&information[0])
		err := gl.GoStr(&information[0])
		return packagederror.NewError(packagederror.ProgramLinkFail,err)
	}	
	return nil
}

func (instance *Shaders) GetProgram() uint32 {
	return instance.program
}
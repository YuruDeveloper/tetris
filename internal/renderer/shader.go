package renderer

import (
	"os"

	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/go-gl/gl/v4.6-core/gl"
)

const VertexShaderFile = "./public/vertexShader.vs"
const FragmentShaderFile = "./public/fragmentShader.vs"  

func NewShaders() *Shaders {
	return &Shaders{}
}

type Shaders struct{
	program uint32
	vertexShaderSource **uint8
	freeVertexShaderSource func()
	vertexShaderSourceLength int32
	fragmentShaderSource **uint8
	freeFragmentShaderSource func()
	fragmentShaderSourceLength int32
}

func (instance *Shaders)LoadFiles() error {
	if instance.program != 0 && gl.IsProgram(instance.program) {
		return packagederror.NewError(packagederror.AlreadyCompiled,"Already Compiled") 
	}
	data , err := os.ReadFile(VertexShaderFile)
	if err != nil {
		return packagederror.NewError(packagederror.FailReadFile,err.Error())
	}
	instance.vertexShaderSource , instance.freeVertexShaderSource = gl.Strs(string(data))
	instance.vertexShaderSourceLength = int32(len(string(data)))
	data , err = os.ReadFile(FragmentShaderFile)
	if err != nil {
		return packagederror.NewError(packagederror.FailReadFile,err.Error())
	}
	instance.fragmentShaderSource , instance.freeFragmentShaderSource = gl.Strs(string(data))
	instance.fragmentShaderSourceLength = int32(len(string(data)))
	return nil
}

func (instance *Shaders) CompileShaders() error {
	if instance.program != 0 && gl.IsProgram(instance.program) {
		return packagederror.NewError(packagederror.AlreadyCompiled,"Already Compiled") 
	}
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vertexShader,1,instance.vertexShaderSource,&instance.vertexShaderSourceLength)
	gl.CompileShader(vertexShader)
	err := instance.checkShaderStatus(vertexShader)
	if err != nil {
		return err
	}
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(fragmentShader,1,instance.fragmentShaderSource,&instance.fragmentShaderSourceLength)
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
		instance.freeVertexShaderSource()
		instance.freeFragmentShaderSource()
	}
	return err
}

func (instance *Shaders) checkShaderStatus(Shader uint32) error {
	var isCompile int32
	gl.GetShaderiv(Shader,gl.COMPILE_STATUS,&isCompile)
	if isCompile == gl.FALSE {
		var maxLength int32
		gl.GetShaderiv(Shader,gl.INFO_LOG_LENGTH,&maxLength)
		if maxLength == 0 {
			return packagederror.NewError(packagederror.ShaderCompileFail,"No Logs")
		}
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
		if maxLength == 0 {
			return packagederror.NewError(packagederror.ProgramLinkFail,"No Logs")
		}
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

func (instance *Shaders) CleanProgram() {
	if instance.program == 0 || !gl.IsProgram(instance.program) {
		return
	}
	gl.DeleteProgram(instance.program)
}
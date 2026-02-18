package renderer

import (
	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var _ ports.Shader = (*Shader)(nil)

func NewShaders() ports.Shader {
	return &Shader{}
}

type Shader struct {
	program uint32
}

func (instance *Shader) CompileShader(vertexShaderString string,fragmentShaderString string) error {
	if instance.program != 0 && gl.IsProgram(instance.program) {
		return packagederror.NewError(packagederror.AlreadyCompiled, "Already Compiled")
	}

	vertexShaderSourceLength := int32(len(vertexShaderString))
	fragmentShaderSourceLength := int32(len(fragmentShaderString))
	vertexShaderSource , freeVertexShaderSource := gl.Strs(vertexShaderString)
	fragmentShaderSource , freeFragmentShaderSource:= gl.Strs(fragmentShaderString)

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vertexShader, 1, vertexShaderSource, &vertexShaderSourceLength)
	gl.CompileShader(vertexShader)
	err := instance.checkShaderStatus(vertexShader)
	if err != nil {
		freeVertexShaderSource()
		freeFragmentShaderSource()
		return err
	}

	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(fragmentShader, 1, fragmentShaderSource, &fragmentShaderSourceLength)
	gl.CompileShader(fragmentShader)
	err = instance.checkShaderStatus(fragmentShader)
	if err != nil {
		freeVertexShaderSource()
		freeFragmentShaderSource()
		return err
	}

	instance.program = gl.CreateProgram()
	gl.AttachShader(instance.program, vertexShader)
	gl.AttachShader(instance.program, fragmentShader)
	gl.LinkProgram(instance.program)
	
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	freeVertexShaderSource()
	freeFragmentShaderSource()
	
	err = instance.checkProgramStatus(instance.program)
	
	return err
}

func (instance *Shader) checkShaderStatus(Shader uint32) error {
	var isCompile int32
	gl.GetShaderiv(Shader, gl.COMPILE_STATUS, &isCompile)
	if isCompile == gl.FALSE {
		var maxLength int32
		gl.GetShaderiv(Shader, gl.INFO_LOG_LENGTH, &maxLength)
		if maxLength == 0 {
			return packagederror.NewError(packagederror.ShaderCompileFail, "No Logs")
		}
		information := make([]uint8, maxLength)
		gl.GetShaderInfoLog(Shader, maxLength, &maxLength, &information[0])
		err := gl.GoStr(&information[0])
		return packagederror.NewError(packagederror.ShaderCompileFail, err)
	}
	return nil
}

func (instance *Shader) checkProgramStatus(program uint32) error {
	var isLinked int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &isLinked)
	if isLinked == gl.FALSE {
		var maxLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &maxLength)
		if maxLength == 0 {
			return packagederror.NewError(packagederror.ProgramLinkFail, "No Logs")
		}
		information := make([]uint8, maxLength)
		gl.GetProgramInfoLog(program, maxLength, &maxLength, &information[0])
		err := gl.GoStr(&information[0])
		return packagederror.NewError(packagederror.ProgramLinkFail, err)
	}
	return nil
}

func (instance *Shader) GetProgram() types.Program {
	return types.Program(instance.program)
}

func (instance *Shader) Delete() {
	if instance.program == 0 || !gl.IsProgram(instance.program) {
		return
	}
	gl.DeleteProgram(instance.program)
}

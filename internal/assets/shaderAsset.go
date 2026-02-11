package asset

import (
	"os"

	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var _ ports.Asset = (*ShaderAsset)(nil)

type ShaderAsset struct {
	vertexShaderSource         **uint8
	freeVertexShaderSource     func()
	vertexShaderSourceLength   int32
	fragmentShaderSource       **uint8
	freeFragmentShaderSource   func()
	fragmentShaderSourceLength int32
	vertexShaderFile string
	fragmentShaderFile string
	shader ports.Shader
}

func NewShaderAsset(createFunction func() ports.Shader ,vertexShaderFile string,fragmentShaderFile string) *ShaderAsset {
	return &ShaderAsset{
		vertexShaderFile: vertexShaderFile,
		fragmentShaderFile: fragmentShaderFile,
		shader: createFunction(),
	}
}

func (instance *ShaderAsset) Load() error {
	err := instance.loadFiles()
	if err != nil {
		return err
	}
	err = instance.init()
	return err
}

func (instance *ShaderAsset) UnLoad() {
	instance.shader.Delete()
}

func (instance *ShaderAsset) IsLoaded() bool {
	return gl.IsProgram(uint32(instance.shader.GetProgram()))
}

func (instance *ShaderAsset) loadFiles() error {
	if instance.shader.GetProgram() != 0 && gl.IsProgram(uint32(instance.shader.GetProgram())) {
		return packagederror.NewError(packagederror.AlreadyCompiled, "Already Compiled")
	}
	data, err := os.ReadFile(instance.vertexShaderFile)
	if err != nil {
		return packagederror.NewError(packagederror.FailReadFile, err.Error())
	}
	instance.vertexShaderSource, instance.freeVertexShaderSource = gl.Strs(string(data))
	instance.vertexShaderSourceLength = int32(len(string(data)))
	data, err = os.ReadFile(instance.fragmentShaderFile)
	if err != nil {
		return packagederror.NewError(packagederror.FailReadFile, err.Error())
	}
	instance.fragmentShaderSource, instance.freeFragmentShaderSource = gl.Strs(string(data))
	instance.fragmentShaderSourceLength = int32(len(string(data)))
	return nil
}

func (instance *ShaderAsset) init() error {
	if instance.shader.GetProgram() != 0 && gl.IsProgram(uint32(instance.shader.GetProgram())) {
		return packagederror.NewError(packagederror.AlreadyCompiled, "Already Compiled")
	}
	err := instance.shader.CompileShader(instance.vertexShaderSource,instance.vertexShaderSourceLength,instance.fragmentShaderSource,instance.fragmentShaderSourceLength)
	if err != nil {
		return err
	}
	instance.freeVertexShaderSource()
	instance.freeFragmentShaderSource()
	return nil
}

func (instance *ShaderAsset) Get() types.Program {
	return instance.shader.GetProgram()
}
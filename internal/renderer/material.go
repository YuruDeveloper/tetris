package renderer

import (
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

type Material struct {
	program *types.Reference[types.Program]
	texture *types.Reference[types.Texture]
}

func NewMeterial(program *types.Reference[types.Program],texture *types.Reference[types.Texture]) *Material {
	return &Material{
		program: program,
		texture: texture,
	}
}

func (instance *Material) Init() {
	gl.UseProgram(uint32(instance.program.Get()))
	samplerLocation := gl.GetUniformLocation(uint32(instance.program.Get()),gl.Str("textureMap\x00"))
    gl.Uniform1i(samplerLocation, 0)
}

func (instance *Material) GetProgram() types.Program {
	return instance.program.Get()
}

func (instance *Material) Render() {
	gl.UseProgram(uint32(instance.program.Get()))
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D,uint32(instance.texture.Get()))
}

func (instance *Material) Delete() {
	instance.program.Delete()
	instance.texture.Delete()
}
 
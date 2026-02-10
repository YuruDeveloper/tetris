package renderer

import (
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

type Material struct {
	program types.Program
	texture types.Texture
}

func NewMeterial(program types.Program,texture types.Texture) *Material {
	return &Material{
		program: program,
		texture: texture,
	}
}

func (instance *Material) Init() {
	gl.UseProgram(uint32(instance.program))
	samplerLocation := gl.GetUniformLocation(uint32(instance.program),gl.Str("textureMap\x00"))
    gl.Uniform1i(samplerLocation, 0)
}

func (instance *Material) GetProgram() types.Program {
	return instance.program
}

func (instance *Material) Render() {
	gl.UseProgram(uint32(instance.program))
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D,uint32(instance.texture))
}

func (instance *Material) Delete() {
	
}
 
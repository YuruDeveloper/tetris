package renderer

import (
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

type Material struct {
	program *types.Reference[types.Program]
	texture *types.Reference[types.Texture]
}

func NewMaterial(program *types.Reference[types.Program],texture *types.Reference[types.Texture]) ports.Material {
	return &Material{
		program: program,
		texture: texture,
	}
}

func (instance *Material) Init() {
	program := instance.program.Get()
	gl.UseProgram(uint32(program))
	samplerLocation := gl.GetUniformLocation(uint32(program),gl.Str("textureMap\x00"))
    gl.Uniform1i(samplerLocation, 0) 
}

func (instance *Material) GetMeterial() types.Meterial {
	program  := instance.program.Get()
	texture:= instance.texture.Get()
	return types.Meterial{
		Program: program,	
		Texture: texture,
	}
}

func (instance *Material) Delete() {
	instance.program.Delete()
	instance.texture.Delete()
}
 
package renderer

import "github.com/YuruDeveloper/tetris/internal/types"

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
	
}

package renderer

type Material struct {
	program uint32
	texture *Texture
}

func NewMeterial(program uint32,texture *Texture) *Material {
	return &Material{
		program: program,
		texture: texture,
	}
}

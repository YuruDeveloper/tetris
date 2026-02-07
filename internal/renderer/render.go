package renderer

import (
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var _ ports.Renderer = (*Renderer)(nil)

func NewRenderer(program uint32) *Renderer {
	return &Renderer{
		program: program,
	}
}

type Renderer struct {
	program uint32
}

func (instance *Renderer) Rendering() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.UseProgram(instance.program)
}

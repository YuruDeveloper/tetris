package renderer

import (
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/google/uuid"
)

var _ ports.Renderer = (*Renderer)(nil)

func NewRenderer(program uint32) *Renderer {
	return &Renderer{
		program: program,
		datas:   make(map[uuid.UUID]*RenderingData[types.Vector2]),
	}
}

type Renderer struct {
	program uint32
	datas   map[uuid.UUID]*RenderingData[types.Vector2]
}

func (instance *Renderer) Rendering(deltaTime float64) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.UseProgram(instance.program)
	for _, data := range instance.datas {
		data.Rendering(instance.program)
	}
}

func (instance *Renderer) Delete() {
	for _, data := range instance.datas {
		data.Delete()
	}
}

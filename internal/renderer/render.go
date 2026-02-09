package renderer

import (
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/google/uuid"
)

var _ ports.Renderer = (*Renderer)(nil)

func NewRenderer() *Renderer {
	return &Renderer{
		datas: make(map[uuid.UUID]*RenderObject[types.Vector2]),
	}
}

type Renderer struct {
	datas     map[uuid.UUID]*RenderObject[types.Vector2]
	locationX float64
}

func (instance *Renderer) Set(data *RenderObject[types.Vector2]) {
	instance.datas[uuid.New()] = data
}

func (instance *Renderer) Rendering(deltaTime float64) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	for _, data := range instance.datas {
		instance.locationX += deltaTime * 0.1
		if instance.locationX >= 1 {
			instance.locationX = -1
		}
		data.SetLocation(types.NewVector2(float32(instance.locationX), 0))
		data.Rendering()
	}
}

func (instance *Renderer) Delete() {
	for _, data := range instance.datas {
		data.Delete()
	}
}

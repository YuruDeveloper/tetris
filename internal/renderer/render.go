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
		sync:      NewSync(),
	}
}

type Renderer struct {
	datas     map[uuid.UUID]*RenderObject[types.Vector2]
	sync           *Sync
	locationX float64
}

func (instance *Renderer) Set(data *RenderObject[types.Vector2]) {
	instance.datas[uuid.New()] = data
}

func (instance *Renderer) Rendering(deltaTime float64) {
	instance.sync.WaitSync()
	gl.Clear(gl.COLOR_BUFFER_BIT)
	for _, data := range instance.datas {
		instance.locationX += deltaTime * 100
		if instance.locationX >= 250 {
			instance.locationX = -250
		}
		data.SetLocation(types.NewVector2(float32(instance.locationX), 0))
		data.Rendering()
	}
	instance.sync.NewFence()
}

func (instance *Renderer) Delete() {
	for _, data := range instance.datas {
		data.Delete()
	}
	instance.sync.Delete()
}

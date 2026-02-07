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
		datas: make(map[uuid.UUID]*RenderingData),
		locationX: -1,
	}
}

type Renderer struct {
	program uint32
	datas map[uuid.UUID]*RenderingData
	locationX float32
}

func (instance *Renderer) Init() {
	vertices := []types.Vector2 {
		types.NewVector2(0.2,0.1),
		types.NewVector2(0.2,-0.1),
		types.NewVector2(-0.2,-0.1),
		types.NewVector2(-0.2,0.1),
	}
	indices := []uint32 {
		0 ,1 ,2,
		0, 2, 3,
	}
	triID := uuid.New()

	instance.datas[triID] = NewRenderingData()
	instance.datas[triID].InitData(vertices,indices,types.NewVector2(1,1),types.NewVector2(0,0))
}


func (instance *Renderer) Rendering(deltaTime float64) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.UseProgram(instance.program)
	for _ , data := range instance.datas {
			instance.locationX += float32(0.1 * deltaTime)
		if instance.locationX >= 1 {
			instance.locationX = -1
		}
		data.SetLocation(types.NewVector2(instance.locationX,0))
		data.Rendering(instance.program)
	}
}

func (instance *Renderer) ClearDatas() {
	for _ ,data := range instance.datas {
		data.Delete()
	}
}

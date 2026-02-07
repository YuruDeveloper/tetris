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
	}
}

type Renderer struct {
	program uint32
	datas map[uuid.UUID]*RenderingData
}

func (instance *Renderer) Init() {
	vertices := []types.Vector2 {
		types.NewVector2(-0.2,-0.1),
		types.NewVector2(0.2,-0.1),
		types.NewVector2(0.0,0.1),
	}
	triID := uuid.New()

	instance.datas[triID] = NewRenderingData()
	instance.datas[triID].InitData(vertices,types.NewVector2(1,1),types.NewVector2(0,0))
	gl.BindVertexArray(instance.datas[triID].GetVertexArrayObject())
}


func (instance *Renderer) Rendering() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.UseProgram(instance.program)
	gl.DrawArrays(gl.TRIANGLES,0,3)
}

func (instance *Renderer) ClearDatas() {
	for _ ,data := range instance.datas {
		data.Delete()
	}
}

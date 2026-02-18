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
		datas: make(map[uuid.UUID]*RenderObject),
		idList: make(map[uuid.UUID]int),
		sync:      NewSync(),	
	}
}

type Renderer struct {
	datas     map[uuid.UUID]*RenderObject
	idList map[uuid.UUID]int
	transform *Transform2D
	sync           *Sync
	world *World
}

func (instance *Renderer) Init(viewport types.Vector2) {
	instance.world = NewWorldTransform(viewport)
	instance.transform = NewTransform()
}

func (instance *Renderer) NewObject(mesh *types.Reference[types.Mesh],material *types.Handle[types.Meterial],location,size types.Vector2) uuid.UUID {
	uuid := uuid.New()
	id :=len(instance.datas)
	instance.idList[uuid] = id
	instance.transform.NewTransform(id,types.PackedTransform[types.Vector2]{ Size: size,Location : location })
	instance.datas[uuid] = NewRenderObject(uint32(id),mesh,material,instance.world,instance.transform)
	return uuid
}

func (instance *Renderer) SetSize(uuid uuid.UUID,size types.Vector2) {
	instance.transform.SetSize(instance.idList[uuid],size)
}

func (instance *Renderer) SetLocation(uuid uuid.UUID,location types.Vector2) {
	instance.transform.SetLocation(instance.idList[uuid],location)
}

func (instance *Renderer) Rendering(deltaTime float64) {
	instance.sync.WaitSync()
	gl.Clear(gl.COLOR_BUFFER_BIT)
	for _, data := range instance.datas {
		data.Rendering()
	}
	instance.sync.NewFence()
}

func (instance *Renderer) Delete() {
	for _, data := range instance.datas {
		data.Delete()
	}
	instance.world.Delete()
	instance.sync.Delete()
}

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
		objects: make([][]*RenderObject, 0),
		idList: make(map[uuid.UUID]int),
		sync:      NewSync(),	
	}
}

type Renderer struct {
	objects [][]*RenderObject
	idList map[uuid.UUID]int
	transform *Transform2D
	sync           *Sync
	world *World
}

func (instance *Renderer) Init(viewport types.Vector2) {
	instance.world = NewWorldTransform(viewport)
	instance.world.Init()
	instance.transform = NewTransform()
}

func (instance *Renderer) NewObject(order int,mesh *types.Reference[types.Mesh],material *types.Handle[types.Meterial],location,size types.Vector2) uuid.UUID {
	for order >= len(instance.objects) {
		instance.objects = append(instance.objects, make([]*RenderObject,0))
	}
	uuid := uuid.New()
	id := 0
	for _ , slice := range instance.objects {
		id += len(slice) 
	}
	instance.idList[uuid] = id
	instance.transform.NewTransform(id,types.PackedTransform[types.Vector2]{ Size: size,Location : location })
	instance.objects[order] = append(instance.objects[order], NewRenderObject(uint32(id),mesh,material,instance.world,instance.transform) )
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
	for _ , slice := range instance.objects {
		for _ , object := range slice {
			object.Rendering()
		}
	}
	instance.sync.NewFence()
}

func (instance *Renderer) Delete() {
	for _ , slice := range instance.objects {
		for _ , object := range slice {
			object.Delete()
		}
	}
	instance.transform.Delete()
	instance.world.Delete()
	instance.sync.Delete()
}

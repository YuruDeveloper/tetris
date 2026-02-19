package renderer

import (
	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/google/uuid"
)

var _ ports.Renderer = (*Renderer)(nil)

func NewRenderer(manager ports.Manager) *Renderer {
	return &Renderer{
		objects: make([]map[uuid.UUID]*RenderObject, 0),
		idList: make(map[uuid.UUID]int),
		sync:      NewSync(),	
		idManager: NewIDManager(),
		assetManager: manager,
	}
}

type Renderer struct {
	objects []map[uuid.UUID]*RenderObject
	idList map[uuid.UUID]int
	assetManager ports.Manager
	idManager *IDManager
	transform *Transform2D
	sync           *Sync
	world *World
}

func (instance *Renderer) Init(viewport types.Vector2) {
	instance.world = NewWorldTransform(viewport)
	instance.world.Init()
	instance.transform = NewTransform()
}

func (instance *Renderer) NewObject(order int,meshID uuid.UUID,materialId uuid.UUID ,location,size types.Vector2) (uuid.UUID , error) {
	mesh ,err := instance.assetManager.MeshAsset2D(meshID)
	if err != nil {
		return uuid.Nil , packagederror.NewError(packagederror.FailLoadAsset,err.Error())
	}
	material , err := instance.assetManager.Material(materialId)
	if err != nil {
		mesh.Delete()
		return uuid.Nil , packagederror.NewError(packagederror.FailLoadAsset,err.Error())
	}
	for order >= len(instance.objects) {
		instance.objects = append(instance.objects,make(map[uuid.UUID]*RenderObject))
	}
	uuid := uuid.New()
	id := instance.idManager.Get()
	instance.idList[uuid] = id
	instance.transform.NewTransform(id,types.PackedTransform[types.Vector2]{ Size: size,Location : location })
	instance.objects[order][uuid] = NewRenderObject(uint32(id),mesh,material,instance.world,instance.transform)
	return uuid , nil
}

func (instance *Renderer) DeleteObject(uuid uuid.UUID) {
	for _ , objectMap := range instance.objects {
		if object , ok := objectMap[uuid] ; ok {
			object.Delete()
			instance.idManager.Delete(instance.idList[uuid])
			delete(instance.idList,uuid)
			delete(objectMap,uuid)
		}
	}
}

func (instance *Renderer) SetSize(uuid uuid.UUID,size types.Vector2) {
	if id , ok := instance.idList[uuid] ; ok {
		instance.transform.SetSize(id,size)
	}
}

func (instance *Renderer) SetLocation(uuid uuid.UUID,location types.Vector2) {
	if id , ok := instance.idList[uuid] ; ok {
		instance.transform.SetLocation(id,location)
	}
}

func (instance *Renderer) Rendering(deltaTime float64) {
	instance.sync.WaitSync()
	gl.Clear(gl.COLOR_BUFFER_BIT)
	for _ , objectMap := range instance.objects {
		for _ , object := range objectMap {
			object.Rendering()
		}
	}
	instance.sync.NewFence()
}

func (instance *Renderer) Delete() {
	for _ , objectMap := range instance.objects {
		for _ , object := range objectMap {
			object.Delete()
		}
	}
	instance.transform.Delete()
	instance.world.Delete()
	instance.sync.Delete()
}

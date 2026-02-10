package renderer

import (
	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/types"
)

func NewRenderObject[T types.Vector](vertex []T, indices []uint32,uv []types.Vector2, size T, location T,viewportSize types.Vector2,material *Material) (*RenderObject[T], error) {
	if len(vertex) == 0 || len(indices) == 0  {
		return nil, packagederror.NewError(packagederror.FailCreateRenderingData, "Something is empty or nil")
	}
	mesh, err := NewMesh(vertex, indices,uv)
	if err != nil {
		return nil, err
	}
	transform := NewTransform(size, location,viewportSize)
	return &RenderObject[T]{
		mesh:      mesh,
		material: material,
		transform: transform,
		sync:      NewSync(),
	}, nil
}

type RenderObject[T types.Vector] struct {
	mesh           *Mesh[T]
	transform      *Transform[T]
	material *Material
	sync           *Sync
}

func (instance *RenderObject[T]) Init(transformIndex uint32) error {
	instance.material.Init()
	err := instance.mesh.Init()
	if err != nil {
		instance.Delete()
		return err
	}
	instance.transform.Init(instance.material.GetProgram()) 
	return nil
}

func (instance *RenderObject[T]) SetLocation(location T) {
	instance.transform.SetLocation(location)
}

func (instance *RenderObject[T]) SetSize(size T) {
	instance.transform.SetSize(size)
}

func (instance *RenderObject[T]) Rendering() {
	instance.sync.WaitSync()
	instance.material.Render()
	instance.mesh.Render()
	instance.sync.NewFence()
}

func (instance *RenderObject[T]) Delete() {
	instance.material.Delete()
	instance.mesh.Delete()
	instance.transform.Delete()
	instance.sync.Delete()
}

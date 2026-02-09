package renderer

import (
	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

func NewRenderingData[T types.Vector](vertex []T, indices []uint32, size T, location T,program types.Program, texture types.Texture) (*RenderingData[T], error) {
	if len(vertex) == 0 || len(indices) == 0  {
		return nil, packagederror.NewError(packagederror.FailCreateRenderingData, "Something is empty or nil")
	}
	mesh, err := NewMesh(vertex, indices)
	if err != nil {
		return nil, err
	}
	transform := NewTransform(size, location)
	return &RenderingData[T]{
		mesh:      mesh,
		texture:   texture,
		shader: program,
		transform: transform,
		sync:      NewSync(),
	}, nil
}

type RenderingData[T types.Vector] struct {
	mesh           *Mesh[T]
	transform      *Transform[T]
	texture        types.Texture
	shader types.Program
	programIndex   uint32
	transformIndex uint32
	sync           *Sync
}

func (instance *RenderingData[T]) Init(transformIndex uint32) error {
	instance.transformIndex = transformIndex
	instance.programIndex = gl.GetUniformBlockIndex(uint32(instance.shader), gl.Str("TransformBlock\x00"))
	err := instance.mesh.Init()
	if err != nil {
		instance.Delete()
		return err
	}
	gl.BindBufferBase(gl.UNIFORM_BUFFER, instance.transformIndex, instance.transform.GetBuffer())
	instance.transform.Binding(uint32(instance.shader), instance.programIndex, instance.transformIndex)
	return nil
}

func (instance *RenderingData[T]) SetLocation(location T) {
	instance.transform.SetLocation(location)
}

func (instance *RenderingData[T]) SetSize(size T) {
	instance.transform.SetSize(size)
}

func (instance *RenderingData[T]) Rendering() {
	instance.sync.WaitSync()
	gl.UseProgram(uint32(instance.shader))
	instance.mesh.Rendering()
	instance.sync.NewFence()
}

func (instance *RenderingData[T]) Delete() {
	instance.mesh.Delete()
	instance.transform.Delete()
	instance.sync.Delete()
}

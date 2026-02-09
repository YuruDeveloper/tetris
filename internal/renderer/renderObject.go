package renderer

import (
	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

func NewRenderObject[T types.Vector](vertex []T, indices []uint32,uv []types.Vector2, size T, location T,program types.Program, texture types.Texture) (*RenderObject[T], error) {
	if len(vertex) == 0 || len(indices) == 0  {
		return nil, packagederror.NewError(packagederror.FailCreateRenderingData, "Something is empty or nil")
	}
	mesh, err := NewMesh(vertex, indices,uv)
	if err != nil {
		return nil, err
	}
	transform := NewTransform(size, location)
	return &RenderObject[T]{
		mesh:      mesh,
		texture:   texture,
		shader: program,
		transform: transform,
		sync:      NewSync(),
	}, nil
}

type RenderObject[T types.Vector] struct {
	mesh           *Mesh[T]
	transform      *Transform[T]
	texture        types.Texture
	shader types.Program
	programIndex   uint32
	transformIndex uint32
	sync           *Sync
}

func (instance *RenderObject[T]) Init(transformIndex uint32) error {
	instance.transformIndex = transformIndex
	instance.programIndex = gl.GetUniformBlockIndex(uint32(instance.shader), gl.Str("TransformBlock\x00"))
	err := instance.mesh.Init()
	if err != nil {
		instance.Delete()
		return err
	}
	gl.BindBufferBase(gl.UNIFORM_BUFFER, instance.transformIndex, instance.transform.GetBuffer())
	gl.UseProgram(uint32(instance.shader))
	samplerLocation := gl.GetUniformLocation(uint32(instance.shader),gl.Str("textureMap\x00"))
    gl.Uniform1i(samplerLocation, 0)
	instance.transform.Binding(uint32(instance.shader), instance.programIndex, instance.transformIndex)
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
	gl.UseProgram(uint32(instance.shader))
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D,uint32(instance.texture))
	instance.mesh.Rendering()
	instance.sync.NewFence()
}

func (instance *RenderObject[T]) Delete() {
	instance.mesh.Delete()
	instance.transform.Delete()
	instance.sync.Delete()
}

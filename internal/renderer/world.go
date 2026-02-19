package renderer

import (
	"unsafe"

	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

const WorldIndex = uint32(0)

type World struct {
	buffer *Buffer[types.PackedWorldTransform]
	world unsafe.Pointer
}

func NewWorldTransform(viewportSize types.Vector2) *World {
	var zero types.PackedWorldTransform
	flags := gl.MAP_WRITE_BIT | gl.MAP_PERSISTENT_BIT | gl.MAP_COHERENT_BIT
	buffer := NewBufferWithData(types.PackedWorldTransform { ViewportSize: viewportSize },uint32(flags))
	world := gl.MapNamedBufferRange(buffer.GetDataBuffer(),0,int(unsafe.Sizeof(zero)),uint32(flags))
	return &World{
		buffer: buffer,
		world: world,
	}
}

func (instance *World) Init() {
	gl.BindBufferBase(gl.UNIFORM_BUFFER,WorldIndex,instance.buffer.GetDataBuffer())
}

func (instance *World) Bind(program types.Program) {
	index := gl.GetUniformBlockIndex(uint32(program),gl.Str("WorldBlock\x00"))
	gl.UniformBlockBinding(uint32(program),index,WorldIndex)
}

func (instance *World) SetViewportSize(size types.Vector2) {
	worldPointer  := (*types.PackedWorldTransform)(instance.world)
	worldPointer.ViewportSize = size
}

func (instance *World) Delete() {
	gl.UnmapNamedBuffer(instance.buffer.GetDataBuffer())
	instance.buffer.Delete()
}

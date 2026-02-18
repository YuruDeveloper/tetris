package renderer

import (
	"unsafe"

	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

const WorldIndex = uint32(0)

type World struct {
	buffer *Buffer[types.PackedWorldTrasnform]
	world unsafe.Pointer
}

func NewWorldTransform(viewportSize types.Vector2) *World {
	var zero types.PackedWorldTrasnform
	flags := gl.MAP_WRITE_BIT | gl.MAP_PERSISTENT_BIT | gl.MAP_COHERENT_BIT
	buffer := NewBufferWithData(types.PackedWorldTrasnform { ViewportSize: viewportSize },uint32(flags))
	world := gl.MapNamedBufferRange(buffer.GetDataBuffer(),0,int(unsafe.Sizeof(zero)),uint32(flags))
	return &World{
		buffer: buffer,
		world: world,
	}
}

func (instance *World) Init(program types.Program) {
	gl.BindBufferBase(gl.UNIFORM_BUFFER,WorldIndex,instance.buffer.GetDataBuffer())
	gl.UniformBlockBinding(uint32(program),WorldIndex,WorldIndex)
}

func (instance *World) SetViewportSize(size types.Vector2) {
	worldPointer  := (*types.PackedWorldTrasnform)(instance.world)
	worldPointer.ViewportSize = size
}

func (instance *World) Delete() {
	gl.UnmapNamedBuffer(instance.buffer.GetDataBuffer())
	instance.buffer.Delete()
}
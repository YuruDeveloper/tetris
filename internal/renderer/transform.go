package renderer

import (
	"unsafe"

	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

type Transform[T types.Vector] struct {
	transformBuffer uint32
	transform       unsafe.Pointer
}

func NewTransform[T types.Vector](size T, location T,viewportSize types.Vector2) *Transform[T] {
	var transformBuffer uint32
	flag := gl.MAP_WRITE_BIT | gl.MAP_PERSISTENT_BIT | gl.MAP_COHERENT_BIT
	gl.CreateBuffers(1, &transformBuffer)
	gl.NamedBufferStorage(transformBuffer, FloatDataSize*6, unsafe.Pointer(&types.PackedTransform[T]{Size: size, Location: location , ViewportSize : viewportSize }), uint32(flag))
	transform := gl.MapNamedBufferRange(transformBuffer, 0, FloatDataSize*6, uint32(flag))

	return &Transform[T]{
		transformBuffer: transformBuffer,
		transform:       transform,
	}
}

func (instance *Transform[T]) Binding(program, programIndex, transformIndex uint32) {
	gl.UniformBlockBinding(program, programIndex, transformIndex)
}

func (instance *Transform[T]) GetBuffer() uint32 {
	return instance.transformBuffer
}

func (instance *Transform[T]) SetLocation(location T) {
	transformPointer := (*types.PackedTransform[T])(instance.transform)
	transformPointer.Location = location
}

func (instance *Transform[T]) SetSize(size T) {
	transformPointer := (*types.PackedTransform[T])(instance.transform)
	transformPointer.Size = size
}

func (instance *Transform[T]) Delete() {
	gl.UnmapNamedBuffer(instance.transformBuffer)
	gl.DeleteBuffers(1, &instance.transformBuffer)
}

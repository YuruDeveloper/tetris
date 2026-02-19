package renderer

import (
	"unsafe"

	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

const TransformIndex = uint32(1)
const InitSize = 1024

type Transform2D struct {
	buffer *Buffer[types.PackedTransform[types.Vector2]]
	transform []types.PackedTransform[types.Vector2]
	size int
}

func NewTransform() *Transform2D {
	var zero types.PackedTransform[types.Vector2]
	flags := gl.MAP_WRITE_BIT | gl.MAP_PERSISTENT_BIT | gl.MAP_COHERENT_BIT
	buffer , _ := NewBufferWithDatas(make([]types.PackedTransform[types.Vector2],InitSize),uint32(flags))
	transform := gl.MapNamedBufferRange(buffer.GetDataBuffer(), 0, int(unsafe.Sizeof(zero)) * InitSize, uint32(flags))
	transformPointer := (*types.PackedTransform[types.Vector2])(transform)
	listPointer := unsafe.Slice(transformPointer,InitSize)
	gl.BindBufferBase(gl.SHADER_STORAGE_BUFFER,TransformIndex,buffer.GetDataBuffer())
	return &Transform2D{
		buffer: buffer,
		transform : listPointer,
		size: InitSize,
	}
}

func (instance *Transform2D) Bind(program types.Program) {
	index := gl.GetProgramResourceIndex(uint32(program),gl.SHADER_STORAGE_BLOCK,gl.Str("TranformBlock\x00"))
	gl.ShaderStorageBlockBinding(uint32(program),index,TransformIndex)
}

func (instance *Transform2D) resize() {
	var zero types.PackedTransform[types.Vector2]
	oldBuffer := instance.buffer
	// make new buffer
	flags := gl.MAP_WRITE_BIT | gl.MAP_PERSISTENT_BIT | gl.MAP_COHERENT_BIT
	newBuffer , _ := NewBufferWithDatas(make([]types.PackedTransform[types.Vector2],instance.size * 2),uint32(flags))
	// data copy
	gl.CopyBufferSubData(oldBuffer.GetDataBuffer(),newBuffer.GetDataBuffer(),0,0,int(unsafe.Sizeof(zero)) *  instance.size)
	// setup
	transform := gl.MapNamedBufferRange(newBuffer.GetDataBuffer(), 0, int(unsafe.Sizeof(zero)) * instance.size * 2, uint32(flags))
	gl.BindBufferBase(gl.SHADER_STORAGE_BUFFER,TransformIndex,newBuffer.GetDataBuffer())
	transformPointer := (*types.PackedTransform[types.Vector2])(transform)
	listPointer := unsafe.Slice(transformPointer,instance.size * 2)
	// switch
	instance.size *= 2
	instance.transform = listPointer
	instance.buffer = newBuffer
	// destory
	gl.UnmapNamedBuffer(oldBuffer.GetDataBuffer())
	oldBuffer.Delete()
}

func (instance *Transform2D) NewTransform(id int,data types.PackedTransform[types.Vector2]) {
	for id >= instance.size {
		instance.resize()
	}
	instance.transform[id] = data
	gl.MemoryBarrier(gl.SHADER_STORAGE_BARRIER_BIT)
}

func (instance *Transform2D) SetSize(id int,size types.Vector2) {
	instance.transform[id].Size = size
}

func (instance *Transform2D) SetLocation(id int,location types.Vector2) {
	instance.transform[id].Location = location
}

func (instance *Transform2D) Delete() {
	gl.UnmapNamedBuffer(instance.buffer.GetDataBuffer())
	instance.buffer.Delete()
}




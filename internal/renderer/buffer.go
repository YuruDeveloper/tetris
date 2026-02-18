package renderer

import (
	"unsafe"

	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

type Buffer[T types.BufferData] struct {
	dataBuffer uint32
}

func NewBufferWithData[T types.BufferData](data T,flag uint32) *Buffer[T] {
	var dataBuffer uint32
	var zero T
	size := int(unsafe.Sizeof(zero))
	gl.CreateBuffers(1,&dataBuffer)
	gl.NamedBufferStorage(dataBuffer,size,unsafe.Pointer(&data),flag)
	return &Buffer[T]{
		dataBuffer: dataBuffer,
	}
}

func NewBufferWithDatas[T types.BufferData](data []T,flag uint32) (*Buffer[T], error) {
	if len(data) == 0 {
		return nil, packagederror.NewError(packagederror.DataArrayIsEmpty, "Fail to allocate into buffer")
	}
	var dataBuffer uint32
	var zero T
	size := int(unsafe.Sizeof(zero))
	gl.CreateBuffers(1, &dataBuffer)
	gl.NamedBufferStorage(dataBuffer, size*len(data), unsafe.Pointer(&data[0]), flag)
	return &Buffer[T]{
		dataBuffer: dataBuffer,
	}, nil
}

func (instance *Buffer[T]) GetDataBuffer() uint32 {
	return instance.dataBuffer
}

func (instance *Buffer[T]) Delete() {
	gl.DeleteBuffers(1, &instance.dataBuffer)
	instance.dataBuffer = 0
}

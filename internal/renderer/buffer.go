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

func NewBuffer[T types.BufferData](data []T) (*Buffer[T], error) {
	if len(data) == 0 {
		return nil, packagederror.NewError(packagederror.DataArrayIsEmpty, "Fail to allocate into buffer")
	}
	// 변수 선언
	var dataBuffer uint32
	var zero T
	size := int(unsafe.Sizeof(zero))
	// 초기화
	gl.CreateBuffers(1, &dataBuffer)
	// 할당
	gl.NamedBufferStorage(dataBuffer, size*len(data), unsafe.Pointer(&data[0]), 0)
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

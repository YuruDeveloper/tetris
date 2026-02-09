package renderer

import (
	"unsafe"

	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

type Mesh[T types.Vector] struct {
	vertexBuffer      *Buffer[T]
	indicesBuffer     *IndicesBuffer
	uvBuffer *Buffer[types.Vector2]
	vertexArrayObject uint32
}

func NewMesh[T types.Vector](vertex []T, indices []uint32,uv []types.Vector2) (*Mesh[T], error) {
	vertexBuffer, err := NewBuffer(vertex)
	if err != nil {
		return nil, err
	}
	indicesBuffer, err := NewIndicesBuffer(indices)
	if err != nil {
		vertexBuffer.Delete()
		return nil, err
	}
	uvBuffer , err := NewBuffer(uv)
	if err != nil {
		vertexBuffer.Delete()
		indicesBuffer.Delete()
		return nil , err
	}
	var vertexArrayObject uint32
	gl.CreateVertexArrays(1, &vertexArrayObject)
	return &Mesh[T]{
		vertexBuffer:      vertexBuffer,
		indicesBuffer:     indicesBuffer,
		uvBuffer: uvBuffer,
		vertexArrayObject: vertexArrayObject,
	}, nil
}

func (instance *Mesh[T]) Init() error {
	var zero T
	size := int32(unsafe.Sizeof(zero))
	count := size / int32(FloatDataSize)
	if count < 2 || count > 4 {
		return packagederror.NewError(packagederror.UnSupportedDataType, "This Type is not supported")
	}
	// location 열기
	gl.EnableVertexArrayAttrib(instance.vertexArrayObject, VertexBufferObjectIndex)
	gl.EnableVertexArrayAttrib(instance.vertexArrayObject, UVBufferIndex)
	// vao binding and format setup
	gl.VertexArrayVertexBuffer(instance.vertexArrayObject, VertexBufferObjectIndex, instance.vertexBuffer.GetDataBuffer(), 0, size)
	gl.VertexArrayAttribFormat(instance.vertexArrayObject, VertexBufferObjectIndex, count, gl.FLOAT, false, 0)
	// vao binding and format setup
	gl.VertexArrayVertexBuffer(instance.vertexArrayObject,UVBufferIndex,instance.uvBuffer.GetDataBuffer(),0,int32(FloatDataSize * 2))
	gl.VertexArrayAttribFormat(instance.vertexArrayObject,UVBufferIndex,2,gl.FLOAT,false,0)
	//	binding
	gl.VertexArrayAttribBinding(instance.vertexArrayObject, VertexBufferObjectIndex, VertexBufferObjectIndex)
	gl.VertexArrayAttribBinding(instance.vertexArrayObject,UVBufferIndex,UVBufferIndex)
	// 
	gl.VertexArrayElementBuffer(instance.vertexArrayObject, instance.indicesBuffer.GetDataBuffer())

	gl.VertexArrayBindingDivisor(instance.vertexArrayObject, VertexBufferObjectIndex, 0)
	return nil
}

func (instance *Mesh[T]) GetVertexArrayObject() uint32 {
	return instance.vertexArrayObject
}

func (instance *Mesh[T]) Delete() {
	instance.indicesBuffer.Delete()
	instance.vertexBuffer.Delete()
	gl.DeleteVertexArrays(1, &instance.vertexArrayObject)
}

func (instance *Mesh[T]) Rendering() {
	gl.BindVertexArray(instance.vertexArrayObject)
	gl.DrawElementsWithOffset(gl.TRIANGLES, instance.indicesBuffer.GetIndicesCount(), gl.UNSIGNED_INT, 0)
}

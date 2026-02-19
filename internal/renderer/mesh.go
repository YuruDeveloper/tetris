package renderer

import (
	"unsafe"

	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

const VertexBufferObjectIndex = uint32(0)
const UVBufferIndex = uint32(1)
const FloatDataSize = int(unsafe.Sizeof(float32(0)))

type Mesh[T types.Vector] struct {
	vertexBuffer      *Buffer[T]
	indicesBuffer     *IndicesBuffer
	uvBuffer *Buffer[types.Vector2]
	vertexArrayObject uint32
}

func NewMesh[T types.Vector](vertex []T, indices []uint32,uv []types.Vector2) (ports.Mesh[T], error) {
	vertexBuffer, err := NewBufferWithDatas(vertex,0)
	if err != nil {
		return nil, err
	}
	indicesBuffer, err := NewIndicesBuffer(indices)
	if err != nil {
		vertexBuffer.Delete()
		return nil, err
	}
	uvBuffer , err := NewBufferWithDatas(uv,0)
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
	gl.EnableVertexArrayAttrib(instance.vertexArrayObject, VertexBufferObjectIndex)
	gl.EnableVertexArrayAttrib(instance.vertexArrayObject, UVBufferIndex)

	gl.VertexArrayVertexBuffer(instance.vertexArrayObject, VertexBufferObjectIndex, instance.vertexBuffer.GetDataBuffer(), 0, size)
	gl.VertexArrayAttribFormat(instance.vertexArrayObject, VertexBufferObjectIndex, count, gl.FLOAT, false, 0)

	gl.VertexArrayVertexBuffer(instance.vertexArrayObject, UVBufferIndex, instance.uvBuffer.GetDataBuffer(), 0, int32(FloatDataSize*2))
	gl.VertexArrayAttribFormat(instance.vertexArrayObject, UVBufferIndex, 2, gl.FLOAT, false, 0)

	gl.VertexArrayAttribBinding(instance.vertexArrayObject, VertexBufferObjectIndex, VertexBufferObjectIndex)
	gl.VertexArrayAttribBinding(instance.vertexArrayObject, UVBufferIndex, UVBufferIndex)

	gl.VertexArrayElementBuffer(instance.vertexArrayObject, instance.indicesBuffer.GetDataBuffer())

	gl.VertexArrayBindingDivisor(instance.vertexArrayObject, VertexBufferObjectIndex, 0)
	return nil
}

func (instance *Mesh[T]) GetMesh() types.Mesh {
	return types.Mesh {
		VertexArrayObject: types.VertexArrayObject(instance.vertexArrayObject),
		IndicesCount: instance.indicesBuffer.GetIndicesCount(),
	}
}

func (instance *Mesh[T]) Delete() {
	instance.indicesBuffer.Delete()
	instance.uvBuffer.Delete()
	instance.vertexBuffer.Delete()
	gl.DeleteVertexArrays(1, &instance.vertexArrayObject)
}

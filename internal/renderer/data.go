package renderer

import (
	"unsafe"

	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

const FloatDataSize = int(unsafe.Sizeof(float32(0)))

func NewRenderingData() *RenderingData {
	var shapeBuffer uint32
	var vertexArrayObject uint32
	gl.CreateBuffers(1,&shapeBuffer)
	gl.CreateVertexArrays(1,&vertexArrayObject)
	return &RenderingData{
		shapeBuffer: shapeBuffer,
		vertexArrayObject: vertexArrayObject,
	}
}

type RenderingData struct {
	shapeBuffer uint32
	vertexArrayObject uint32
}

func (instance *RenderingData) InitData(shape []types.Vector2,size types.Vector2,location types.Vector2) {
	// vbo 에 데이터 바인딩
	gl.NamedBufferData(instance.shapeBuffer,FloatDataSize * 2 * len(shape),unsafe.Pointer(&shape[0]),gl.STATIC_DRAW)
	
	// location 열기
	gl.EnableVertexArrayAttrib(instance.vertexArrayObject,0)
	// vao 에 모양 vbo 바인딩
	gl.VertexArrayVertexBuffer(instance.vertexArrayObject,0,instance.shapeBuffer,0,int32(FloatDataSize) * 2)
	gl.VertexArrayAttribFormat(instance.vertexArrayObject,0,2,gl.FLOAT,false,0)
	// 실제 바인딩
	gl.VertexArrayAttribBinding(instance.vertexArrayObject,0,0)
	// 모양 데이터는 한번만 넘기기
	gl.VertexArrayBindingDivisor(instance.vertexArrayObject,0,0)
}



func (instance *RenderingData) GetVertexArrayObject() uint32 {
	return instance.vertexArrayObject
}

func (instance *RenderingData) Delete() {
	gl.DeleteVertexArrays(1,&instance.vertexArrayObject)
	gl.DeleteBuffers(1,&instance.shapeBuffer)
}
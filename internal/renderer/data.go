package renderer

import (
	"unsafe"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

const FloatDataSize = int(unsafe.Sizeof(float32(0)))
const Uint32DataSize = int(unsafe.Sizeof(uint32(0)))

func NewRenderingData() *RenderingData {
	var shapeBuffer uint32
	var vertexArrayObject uint32
	var indicesBuffer uint32
	var transformBuffer uint32
	
	gl.CreateBuffers(1,&shapeBuffer)
	gl.CreateVertexArrays(1,&vertexArrayObject)
	gl.CreateBuffers(1,&indicesBuffer)
	gl.CreateBuffers(1,&transformBuffer)
	return &RenderingData{
		shapeBuffer: shapeBuffer,
		indicesBuffer: indicesBuffer,
		transformBuffer: transformBuffer,
		vertexArrayObject: vertexArrayObject,
	}
}

type RenderingData struct {
	shapeBuffer uint32
	indicesBuffer uint32
	transformBuffer uint32
	vertexArrayObject uint32
	transoform unsafe.Pointer
	fence uintptr
}

func (instance *RenderingData) InitData(shape []types.Vector2,indices []uint32,size types.Vector2,location types.Vector2) {
	// buffer create
	// 데이터 버퍼에 주입
	gl.NamedBufferStorage(instance.shapeBuffer,FloatDataSize*2*len(shape),unsafe.Pointer(&shape[0]),0)
	gl.NamedBufferStorage(instance.indicesBuffer,Uint32DataSize * len(indices),unsafe.Pointer(&indices[0]),0)
	flag := gl.MAP_WRITE_BIT | gl.MAP_PERSISTENT_BIT | gl.MAP_COHERENT_BIT
	gl.NamedBufferStorage(instance.transformBuffer,FloatDataSize* 4,unsafe.Pointer(&types.PackedTransform{ Size :size , Location: location }),uint32(flag))
	// location 열기
	gl.EnableVertexArrayAttrib(instance.vertexArrayObject,0)
	// vbo 형식 정하기
	gl.VertexArrayVertexBuffer(instance.vertexArrayObject,0,instance.shapeBuffer,0,int32(FloatDataSize) * 2)
	gl.VertexArrayAttribFormat(instance.vertexArrayObject,0,2,gl.FLOAT,false,0)
	// 실제 바인딩
	gl.VertexArrayAttribBinding(instance.vertexArrayObject,0,0)
	gl.VertexArrayElementBuffer(instance.vertexArrayObject,instance.indicesBuffer)
	// mapping
	instance.transoform = gl.MapNamedBufferRange(instance.transformBuffer,0,FloatDataSize * 4 ,uint32(flag))
	gl.BindBufferBase(gl.UNIFORM_BUFFER,0,instance.transformBuffer)
	// 모양 데이터는 한번만 넘기기
	gl.VertexArrayBindingDivisor(instance.vertexArrayObject,0,0)
}

func (instance *RenderingData) SetLocation(location types.Vector2) {
	transoformPointer := (*types.PackedTransform)(instance.transoform)
	transoformPointer.Location = location	
}

func (instance *RenderingData) SetSize(size types.Vector2) {
	transoformPointer := (*types.PackedTransform)(instance.transoform)
	transoformPointer.Size = size	
}

func (instance *RenderingData) WaitSync() {
	if instance.fence != 0 {
		gl.ClientWaitSync(instance.fence,gl.SYNC_FLUSH_COMMANDS_BIT,100)
		gl.DeleteSync(instance.fence)
		instance.fence = 0
	}
}

func (instance *RenderingData) Rendering(program uint32) {
	gl.BindVertexArray(instance.vertexArrayObject)
	index := gl.GetUniformBlockIndex(program,gl.Str("TransformBlock\x00"))
	gl.UniformBlockBinding(program,index,0)
	gl.DrawElementsWithOffset(gl.TRIANGLES,6,gl.UNSIGNED_INT,0)
	instance.fence = gl.FenceSync(gl.SYNC_FLUSH_COMMANDS_BIT,0)
}

func (instance *RenderingData) Delete() {
	gl.DeleteVertexArrays(1,&instance.vertexArrayObject)
	gl.DeleteBuffers(1,&instance.shapeBuffer)
	gl.DeleteBuffers(1,&instance.indicesBuffer)
	gl.UnmapNamedBuffer(instance.transformBuffer)
	gl.DeleteBuffers(1,&instance.transformBuffer)
	if instance.fence != 0 {
		gl.DeleteSync(instance.fence)
	}
}
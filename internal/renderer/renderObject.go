package renderer

import (
	"unsafe"

	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

func NewRenderObject(id uint32,mesh *types.Reference[types.Mesh],material *types.Handle[types.Meterial],world *World,tranfrom *Transform2D) *RenderObject {
	return &RenderObject{
		id : id,
		world: world,
		transform: tranfrom,
		mesh:      mesh,
		material: material,
	}
}

type RenderObject struct {
	id uint32
	transform *Transform2D
	world *World
	mesh     *types.Reference[types.Mesh]
	material *types.Handle[types.Meterial]
}

func (instance *RenderObject) Rendering() {
	mesh := instance.mesh.Get()
	program := instance.material.Get().Program
	instance.material.Render()
	instance.world.Bind(program)
	instance.transform.Bind(program)
	gl.BindVertexArray(uint32(mesh.VertexArrayObject))
	gl.DrawElementsInstancedBaseInstance(gl.TRIANGLES, mesh.IndciesCount, gl.UNSIGNED_INT, unsafe.Pointer(nil), 1, instance.id)
}

func (instance *RenderObject) Delete() {
	instance.material.Delete()
	instance.mesh.Delete()
}

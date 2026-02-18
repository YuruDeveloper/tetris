package asset

import (
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

 var _ ports.Asset = (*MeshAsset2D)(nil)

func New2DMeshAssetWithValues(createFunc func(vertex []types.Vector2, indices []uint32,uv []types.Vector2) (ports.Mesh[types.Vector2], error),vertex []types.Vector2,indices []uint32,uv []types.Vector2) (*MeshAsset2D,error) {
	mesh , err := createFunc(vertex,indices,uv)
	if err != nil {
		return nil , err
	}
	return &MeshAsset2D{
		vertexs: vertex,
		indices: indices,
		uv: uv,
		mesh: mesh,
	} , nil
}


type MeshAsset2D struct {
	vertexs []types.Vector2
	indices []uint32
	uv []types.Vector2
	mesh ports.Mesh[types.Vector2]
}

func (instance *MeshAsset2D) IsLoaded() bool {
	return gl.IsVertexArray(uint32(instance.mesh.GetMesh().VertexArrayObject))
}

func (instance *MeshAsset2D) Load() error {
	return instance.mesh.Init()
}

func (instance *MeshAsset2D) UnLoad() {
	instance.mesh.Delete()
}

func (instance *MeshAsset2D) Get() types.Mesh {
	return instance.mesh.GetMesh()
}


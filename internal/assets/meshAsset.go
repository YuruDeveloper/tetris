package asset

import (
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
)

 var _ ports.Asset = (*MeshAsset2D)(nil)

func New2DMeshAssetWithValues(createFunc func(vertex []types.Vector2, indices []uint32,uv []types.Vector2) (ports.Mesh[types.Vector2], error),vertex []types.Vector2,indices []uint32,uv []types.Vector2) (*MeshAsset2D ,error) {
	mesh , err := createFunc(vertex,indices,uv)
	return &MeshAsset2D{
		vertexs: vertex,
		indices: indices,
		uv: uv,
		mesh: mesh,
		isLoad: false,
	} , err
}


type MeshAsset2D struct {
	vertexs []types.Vector2
	indices []uint32
	uv []types.Vector2
	mesh ports.Mesh[types.Vector2]
	isLoad bool
}

func (instance *MeshAsset2D) IsLoaded() bool {
	return instance.isLoad
}

func (instance *MeshAsset2D) Load() error {
	err := instance.mesh.Init()
	if err == nil {
		instance.isLoad = true
	}
	return err
}

func (instance *MeshAsset2D) UnLoad() {
	instance.isLoad = false
	instance.mesh.Delete()
}

func (instance *MeshAsset2D) Get() types.Mesh {
	return instance.mesh.GetMesh()
}


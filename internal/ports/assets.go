package ports

import (
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/google/uuid"
)

type Asset interface {
	Load() error
	UnLoad()
	IsLoaded() bool
}

type Manager interface {
	Register(uuid uuid.UUID,createFunc func() Asset) error
	Release(uuid uuid.UUID)
	ShaderAsset(uuid uuid.UUID) (*types.Reference[types.Program], error)
	TextureAsset(uuid uuid.UUID) (*types.Reference[types.Texture], error)
	MeshAsset2D(uuid uuid.UUID) (*types.Reference[types.Mesh], error)
	Material(uuid uuid.UUID) (*types.Handle[types.Material] , error)
}

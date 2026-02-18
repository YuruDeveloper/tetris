package ports

import (
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/google/uuid"
)

type Renderer interface {
	Rendering(deltaTime float64)
	NewObject(mesh *types.Reference[types.Mesh],material *types.Handle[types.Meterial],location,size types.Vector2) uuid.UUID
	SetSize(uuid uuid.UUID,size types.Vector2)
	SetLocation(uuid uuid.UUID,location types.Vector2)
}

type Shader interface {
	CompileShader(vertexShaderString string,fragmentShaderString string) error
	GetProgram() types.Program
	Delete()
}

type Texture interface {
	LoadTextureImage(information types.ImageInformation)
	GetTexture() types.Texture
	Delete()
}

type Material interface {
	Init() 
	GetMeterial() types.Meterial
	Delete()
}

type Mesh[T types.Vector] interface {
	Init() error 
	GetMesh() types.Mesh
	Delete()
}

package ports

import (
	"image"

	"github.com/YuruDeveloper/tetris/internal/types"
)

type Renderer interface {
	Rendering(deltaTime float64)
}

type Shader interface {
	CompileShader(vertexShaderSource **uint8,vertexShaderSourceLength int32,fragmentShaderSource **uint8,fragmentShaderSourceLength int32) error
	GetProgram() types.Program
	Delete()
}

type Texture interface {
	LoadTextureImage(image *image.NRGBA, width,height,level int32)
	GetTexture() types.Texture
	Delete()
}

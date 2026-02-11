package resource

import (
	"log"

	asset "github.com/YuruDeveloper/tetris/internal/assets"
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/renderer"
	"github.com/google/uuid"
)

var DefaultShaderID uuid.UUID = uuid.MustParse("c1ad6299-65bb-4570-ba67-48b89a1413cb")
var DefaultTextureID uuid.UUID = uuid.MustParse("e5c08dbc-c288-4090-adfc-5f2a16c4857f")

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Init() {
	must(asset.GetAssetManager().Register(DefaultShaderID,func() ports.Asset {
		return asset.NewShaderAsset(renderer.NewShaders,"./public/vertexShader.vs","./public/fragmentShader.fs")
	}))
	must(asset.GetAssetManager().Register(DefaultTextureID,func() ports.Asset {
		return  asset.NewTextureAsset(renderer.NewTexture,"./public/texture.png")
	}))
}
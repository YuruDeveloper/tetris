package asset

import (
	"log"

	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/google/uuid"
)

var DefaultShaderID uuid.UUID = uuid.MustParse("c1ad6299-65bb-4570-ba67-48b89a1413cb")
var DefaultTextureID uuid.UUID = uuid.MustParse("e5c08dbc-c288-4090-adfc-5f2a16c4857f")

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (instance *AssetManager) Init() {
	must(instance.register(DefaultShaderID,func() ports.Asset {
		return newShaderAsset("./public/vertexShader.vs","./public/fragmentShader.fs")
	}))
	must(instance.register(DefaultTextureID,func() ports.Asset {
		return  newTextureAsset("./public/texture.png")
	}))
}
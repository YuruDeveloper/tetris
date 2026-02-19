package resource

import (
	"log"

	asset "github.com/YuruDeveloper/tetris/internal/assets"
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/renderer"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/google/uuid"
)

var DefaultShaderID uuid.UUID = uuid.MustParse("c1ad6299-65bb-4570-ba67-48b89a1413cb")
var DefaultTextureID uuid.UUID = uuid.MustParse("e5c08dbc-c288-4090-adfc-5f2a16c4857f")
var DefaultMeshID uuid.UUID = uuid.MustParse("b208d251-b6c0-4ff0-a5e7-6e9ddf7cfa6b")
var DefaultMaterialID uuid.UUID = uuid.MustParse("0c5b8443-1647-4a5d-9212-2987140aa44b")
func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Init(manager ports.Manager) {
	must(manager.Register(DefaultShaderID,func() ports.Asset {
		return asset.NewShaderAsset(renderer.NewShaders,"./public/vertexShader.vs","./public/fragmentShader.fs")
	}))
	must(manager.Register(DefaultTextureID,func() ports.Asset {
		return  asset.NewTextureAsset(renderer.NewTexture,"./public/texture.png")
	}))
	must(manager.Register(DefaultMaterialID,func() ports.Asset {
		shader , err := manager.ShaderAsset(DefaultShaderID)
		must(err)
		texture, err := manager.TextureAsset(DefaultTextureID)
		must(err)
		return asset.NewMaterialAsset(renderer.NewMaterial,shader,texture)
	}))
	must(manager.Register(DefaultMeshID,func() ports.Asset {
		asset , err := asset.New2DMeshAssetWithValues(
			renderer.NewMesh,
			[]types.Vector2{
				types.NewVector2(1, 1),
				types.NewVector2(1, -1),
				types.NewVector2(-1, 1),
				types.NewVector2(-1, -1),
			},
			[]uint32{
				0, 1, 2, 2, 3, 1,
			},
			[]types.Vector2{
				types.NewVector2(1.0,0.0),
				types.NewVector2(1.0,1.0),
				types.NewVector2(0.0,0.0),
				types.NewVector2(0.0,1.0),
			},
		)
		must(err)
		return asset
	}))
}
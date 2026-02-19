package asset

import (
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var _ ports.Asset = (*MaterialAsset)(nil)

func NewMaterialAsset(createFunction func(color types.Color,program *types.Reference[types.Program],texture *types.Reference[types.Texture]) ports.Material,color types.Color,shader *types.Reference[types.Program],texture *types.Reference[types.Texture]) *MaterialAsset {
	return &MaterialAsset{
		material: createFunction(color,shader,texture),
		IsLoad: false,
	}
}

type MaterialAsset struct {
	material ports.Material
	IsLoad bool
}

func (instance *MaterialAsset) IsLoaded() bool  {
	return instance.IsLoad
}

func (instance *MaterialAsset) Load() error {
	instance.IsLoad = true
	instance.material.Init()
	return nil
}

func (instance *MaterialAsset) UnLoad() {
	instance.IsLoad = false
	instance.material.Delete()
}

func (instance *MaterialAsset) Get() types.Material {
	return instance.material.GetMaterial()
}


func (instance *MaterialAsset) GetRenderer() func(types.Material) {
	return func(meterial types.Material) {
		gl.UseProgram(uint32(meterial.Program))
		gl.Uniform4f(meterial.ColorLocation,meterial.Color.Red,meterial.Color.Green,meterial.Color.Blue,meterial.Color.Alpha)
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D,uint32(meterial.Texture))
	}
}
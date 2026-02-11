package asset

import (
	"image"
	"image/png"
	"os"

	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
)

var _ ports.Asset = (*TextureAsset)(nil)

func NewTextureAsset(createFunction func() ports.Texture,textureFile string) *TextureAsset {
	return &TextureAsset{
		textureFile: textureFile,
		texture: createFunction(),
		loaded: false,
	} 
}

type TextureAsset struct {
	textureFile string
	texture ports.Texture
	level int32
	loaded bool
}

func (instance *TextureAsset) IsLoaded() bool {
	return instance.loaded
}

func (instance *TextureAsset) Load() error {
	if err := instance.init() ; err != nil {
		return err
	}
	return nil
}

func (instance *TextureAsset) UnLoad() {
	instance.loaded = false
	instance.texture.Delete()
}

func (instance *TextureAsset) init() error {
	file, err := os.Open(instance.textureFile)
	if err != nil {
		return packagederror.NewError(packagederror.FailOpenFile, err.Error())
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return packagederror.NewError(packagederror.FailDecodeImage, err.Error())
	}
	width := int32(img.Bounds().Dx())
	height := int32(img.Bounds().Dy())
	rgba, ok := img.(*image.NRGBA)
	if !ok {
		return packagederror.NewError(packagederror.FailConvertImage, "FailToConvertImage")
	}
	instance.texture.LoadTextureImage(rgba,width,height,instance.level)
	instance.loaded = true
	return nil
}

func (instance *TextureAsset) Get() types.Texture {
	return instance.texture.GetTexture()
}
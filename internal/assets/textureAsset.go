package asset

import (
	"image"
	"image/png"
	"os"

	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/renderer"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var _ ports.Asset = (*TextureAsset)(nil)

func newTextureAsset(textureFile string) *TextureAsset {
	return &TextureAsset{
		textureFile: textureFile,
		texture: renderer.NewTexture(),
	} 
}

type TextureAsset struct {
	textureFile string
	texture ports.Texture
	level int32
}

func (instance *TextureAsset) IsLoaded() bool {
	return gl.IsTexture(uint32(instance.texture.GetTexture()))
}

func (instance *TextureAsset) Load() error {
	if err := instance.init() ; err != nil {
		return err
	}
	return nil
}

func (instance *TextureAsset) UnLoad() {
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
	return nil
}

func (instance *TextureAsset) Get() types.Texture {
	return instance.texture.GetTexture()
}
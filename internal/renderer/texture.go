package renderer

import (
	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/go-gl/gl/v4.6-core/gl"
	"image"
	"image/png"
	"os"
	"unsafe"
)

func NewTexture() *Texture {
	var texture uint32
	gl.CreateTextures(gl.TEXTURE_2D, 1, &texture)
	gl.TextureParameteri(texture, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TextureParameteri(texture, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TextureParameteri(texture, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TextureParameteri(texture, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	return &Texture{
		texture: texture,
	}
}

type Texture struct {
	texture uint32
}

func (instance *Texture) LoadTextureImage(filePath string) error {
	file, err := os.Open(filePath)
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
	gl.TextureStorage2D(instance.texture, 1, gl.RGBA, width, height)
	rgba, ok := img.(*image.NRGBA)
	if !ok {
		return packagederror.NewError(packagederror.FailConvertImage, "FailToConvertImage")
	}
	gl.TextureSubImage2D(instance.texture, 0, 0, 0, width, height, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&rgba.Pix[0]))
	return nil
}

func (instance *Texture) Rendering(vertexArrayObject uint32) {
	gl.BindTexture(vertexArrayObject, instance.texture)
}

func (instance *Texture) DeleteTexture() {
	gl.DeleteTextures(1, &instance.texture)
}

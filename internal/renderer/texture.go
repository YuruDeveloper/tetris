package renderer

import (
	"image"
	"unsafe"

	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var _ ports.Texture = (*Texture)(nil)

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

func (instance *Texture) LoadTextureImage(image *image.NRGBA, width,height,level int32) {
	gl.TextureStorage2D(instance.texture, 1, gl.RGBA, width, height)
	gl.TextureSubImage2D(instance.texture, level, 0, 0, width, height, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&image.Pix[0]))
}

func (instance *Texture) GetTexture() types.Texture {
	return types.Texture(instance.texture)
}

func (instance *Texture) Delete() {
	gl.DeleteTextures(1, &instance.texture)
}

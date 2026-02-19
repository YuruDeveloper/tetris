package types

import (
	"image"
)

type BufferData interface {
	Vector2 | uint32 |  PackedWorldTrasnform | PackedTransform[Vector2] 
}

type Vector interface {
	Vector2
}

type Vector2 struct {
	X, Y float32
}

func NewVector2(x float32, y float32) Vector2 {
	return Vector2{x, y}
}

type PackedTransform[T Vector] struct {
	Size     T
	Location T
}

type PackedWorldTrasnform struct {
	ViewportSize Vector2
}

type ImageInformation struct {
	Image *image.NRGBA
	Width int32
	Height int32
}

type Mesh struct {
	VertexArrayObject VertexArrayObject
	IndciesCount int32
}

type Meterial struct {
	Program Program
	Texture Texture
	Color ColorFloat
	ColorLocation int32
}

type VertexArrayObject uint32
type Program uint32
type Texture uint32
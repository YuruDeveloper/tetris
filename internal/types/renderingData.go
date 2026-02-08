package types

type BufferData interface {
	Vector2 | uint32
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

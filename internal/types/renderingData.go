package types

type Vector2 struct {
 	X float32
	Y float32
}

func NewVector2(x float32,y float32) Vector2 {
	return Vector2{ x , y}
}

type PackedTransform struct {
	Size Vector2
	Location Vector2
}
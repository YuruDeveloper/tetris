package types

type Vector2 [2]float32

func NewVector2(x float32,y float32) Vector2 {
	return Vector2([]float32{ x, y })
}
package types

func NewColor(red uint8,green uint8,blue uint8,alpha uint8) Color {
	return Color{
		Red: red,
		Green: green,
		Blue: blue,
		Alpha: alpha,
	}
}

type Color struct {
	Red uint8
	Green uint8
	Blue uint8
	Alpha uint8
}

type ColorFloat struct {
	Red float32
	Green float32
	Blue float32
	Alpha float32
}

func (instance *Color) GetColor() ColorFloat {
	return ColorFloat{
		Red: float32(instance.Red) / 255,
		Green: float32(instance.Green) / 255,
		Blue: float32(instance.Blue) / 255,
		Alpha: float32(instance.Alpha) / 255,
	}
}
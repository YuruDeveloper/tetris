package types

func NewColor(red float32,green float32,blue float32,alpha float32) Color {
	return Color{
		Red: red,
		Green: green,
		Blue: blue,
		Alpha: alpha,
	}
}

type Color struct {
	Red float32
	Green float32
	Blue float32
	Alpha float32
}
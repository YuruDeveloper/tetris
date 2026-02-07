package ports

type Renderer interface {
	Rendering(deltaTime float64)
}
package ports

type Asset interface {
	Load() error
	UnLoad()
	IsLoaded() bool
}
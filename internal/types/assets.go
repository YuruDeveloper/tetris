package types

type AssetData interface {
	Texture | Program
}

func NewReference[T AssetData](data T ,delete func()) *Reference[T] {
	return &Reference[T]{
		data: data,
		delete: delete,
	}
}

type Reference[T AssetData] struct {
	data T 
	delete func()
}

func (instance *Reference[T]) Get() T {
	return instance.data
}

func (instance *Reference[T]) Delete() {
	instance.delete()
}
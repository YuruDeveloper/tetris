package types

type Handle[T AssetData] struct {
	*Reference[T]
	render func(T)
}

func NewHandle[T AssetData](data T ,delete func(),render func(T)) *Handle[T]{
	reference := NewReference(data,delete)
	return &Handle[T]{
	 	reference,
		render,
	}
}

func (instance *Handle[T]) Render() {
	instance.render(instance.data)
}
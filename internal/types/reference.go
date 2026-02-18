package types

import (
	"sync"
)

func NewReference[T AssetData](data T ,delete func()) *Reference[T] {
	reference :=  &Reference[T]{
		data: data,
		delete: delete,
	}
	return reference
}

type Reference[T AssetData] struct {
	data T 
	delete func()
	once sync.Once
}

func (instance *Reference[T]) Get() T {
	return instance.data
}

func (instance *Reference[T]) Delete() {
	instance.once.Do(func ()  {
		instance.delete()	
	})
}
package types

import "sync"

type TypeSyncMap[Key any,Value any] struct {
	syncMap sync.Map
}

func (instance *TypeSyncMap[Key, Value]) Store(key Key,value Value) {
	instance.syncMap.Store(key,value)
}

func (instance *TypeSyncMap[Key, Value]) Load(key Key) (Value , bool) {
	value , exist := instance.syncMap.Load(key)
	if !exist {
		var zero Value
		return zero , false
	}
	return value.(Value) , true
} 

func (instance *TypeSyncMap[Key, Value]) Delete(key Key) {
	instance.syncMap.Delete(key)
}

func (instance *TypeSyncMap[Key, Value]) Swap(key Key,value Value) {
	instance.syncMap.Swap(key,value)
}

func (instance *TypeSyncMap[Key, Value]) Clear() {
	instance.syncMap.Clear()
}

func (instance *TypeSyncMap[Key, Value]) LoadOrStore(key Key,value Value) (Value,bool) {
	newValue , loaded := instance.syncMap.LoadOrStore(key,value)
	return newValue.(Value) , loaded
}
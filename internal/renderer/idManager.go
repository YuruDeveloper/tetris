package renderer

func NewIDManager() *IDManager {
	return &IDManager{
		freeList: make([]int, 0),
		idCount: 0,
	}
}

type IDManager struct {
	freeList []int
	idCount int
}

func (instance *IDManager) Get() int {
	var id int
	if len(instance.freeList) == 0 {
		id = instance.idCount
		instance.idCount++
		return id
	}
	id = instance.freeList[0]
	instance.freeList = instance.freeList[1:]
	return id
}

func (instance *IDManager) Delete(id int) {
	instance.freeList = append(instance.freeList, id)
}
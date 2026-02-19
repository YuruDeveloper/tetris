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
	last := len(instance.freeList) -1
	id = instance.freeList[last]
	instance.freeList = instance.freeList[:last]
	return id
}

func (instance *IDManager) Delete(id int) {
	instance.freeList = append(instance.freeList, id)
}
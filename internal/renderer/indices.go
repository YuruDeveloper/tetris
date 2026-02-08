package renderer

type IndicesBuffer struct {
	*Buffer[uint32]
	count int32
}

func NewIndicesBuffer(indices []uint32) (*IndicesBuffer, error) {
	buffer, err := NewBuffer(indices)
	if err != nil {
		return nil , err
	}
	count := int32(len(indices))
	return &IndicesBuffer{
		Buffer: buffer,
		count:  count,
	}, nil
}

func (instance *IndicesBuffer) GetIndicesCount() int32 {
	return instance.count
}

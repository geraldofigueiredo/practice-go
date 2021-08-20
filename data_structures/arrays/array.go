package arrays

import (
	"errors"
)

const (
	minCapacity = 16
	kGrowth     = 2
)

type vector struct {
	size     int
	capacity int
	data     []int
}

func NewVector(capacity int) (VectorInterface, error) {
	if capacity < 1 {
		return vector{}, errors.New("invalid capacity")
	}

	vec := vector{
		capacity: minCapacity,
		size:     0,
	}

	for vec.capacity < capacity {
		vec.capacity *= kGrowth
	}

	vec.data = make([]int, vec.capacity)
	return vec, nil
}

func (v vector) Len() int {
	return v.size
}

func (v vector) Capacity() int {
	return v.capacity
}

func (v vector) GetValueAt(index int) (int, error) {
	if err := v.validateIndex(index); err != nil {
		return 0, err
	}
	return v.data[index], nil
}

func (v vector) validateIndex(index int) error {
	if index < 0 || index > v.size-1 {
		return errors.New("invalid index")
	}
	return nil
}

// Aux methods
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

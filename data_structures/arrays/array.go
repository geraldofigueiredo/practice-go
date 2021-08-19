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

	vec := vector{}

	vec.capacity = max(minCapacity, capacity)
	vec.size = 0
	vec.data = make([]int, vec.capacity)

	return vec, nil
}

func (v vector) Len() int {
	return v.size
}

func (v vector) Capacity() int {
	return v.capacity
}

// Aux methods
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
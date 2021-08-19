package arrays

type VectorInterface interface {
	Len() int
	Capacity() int
	GetValueAt(index int) (int, error)
}

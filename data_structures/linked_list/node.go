package linkedlist

// Node Represents the element that makes up the list
type Node[T comparable] struct {
	data     T
	next     *Node[T]
	previous *Node[T]
}

// NewEmptyNode Creates a empty element
func NewEmptyNode[T comparable]() *Node[T] {
	return &Node[T]{}
}

// NewNode Creates an element with a value stored
func NewNode[T comparable](value T, next *Node[T], previous *Node[T]) *Node[T] {
	return &Node[T]{
		data: value,
		next: next,
	}
}

// SetValue Update the data stored in the object
func (element *Node[T]) SetValue(data T) {
	element.data = data
}

// GetValue Returns the data stored in the object
func (element *Node[T]) GetValue() T {
	return element.data
}

// SetNext Update the element pointed
func (element *Node[T]) SetNext(next *Node[T]) {
	element.next = next
}

// SetPrevious Update the element pointed
func (element *Node[T]) SetPrevious(previous *Node[T]) {
	element.previous = previous
}

// GetNext Gets the element pointed
func (element *Node[T]) GetNext() *Node[T] {
	return element.next
}

// GetPrevious Gets the element pointed
func (element *Node[T]) GetPrevious() *Node[T] {
	return element.previous
}

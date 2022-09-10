package linkedlist

// ListElement Represents the element that makes up the list
type ListElement[T comparable] struct {
	data T
	next *ListElement[T]
}

// NewEmptyListElement Creates a empty element
func NewEmptyListElement[T comparable]() *ListElement[T] {
	return &ListElement[T]{}
}

// NewListElement Creates an element with a value stored
func NewListElement[T comparable](value T, next *ListElement[T]) *ListElement[T] {
	return &ListElement[T]{
		data: value,
		next: next,
	}
}

// SetValue Update the data stored in the object
func (element *ListElement[T]) SetValue(data T) {
	element.data = data
}

// GetValue Returns the data stored in the object
func (element *ListElement[T]) GetValue() T {
	return element.data
}

// SetNext Update the element pointed
func (element *ListElement[T]) SetNext(next *ListElement[T]) {
	element.next = next
}

// GetNext Gets the element pointed
func (element *ListElement[T]) GetNext() *ListElement[T] {
	return element.next
}

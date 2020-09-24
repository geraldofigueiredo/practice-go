package linkedlist

// ListElement Represents the element that makes up the list
type ListElement struct {
	value int64
	next  *ListElement
}

// NewEmptyListElement Creates a empty element
func NewEmptyListElement() *ListElement {
	return &ListElement{}
}

// NewListElement Creates an element with a value stored
func NewListElement(value int64, next *ListElement) *ListElement {
	return &ListElement{
		value: value,
		next:  next,
	}
}

// SetValue Update the value stored in the object
func (element *ListElement) SetValue(value int64) {
	element.value = value
}

// GetValue Returns the value stored in the object
func (element *ListElement) GetValue() int64 {
	return element.value
}

// SetNext Update the element pointed
func (element *ListElement) SetNext(next *ListElement) {
	element.next = next
}

// GetNext Gets the element pointed
func (element *ListElement) GetNext() *ListElement {
	return element.next
}

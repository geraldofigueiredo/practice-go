package linkedlist

// LinkedList Represents the interface to access the linkedList elements
type LinkedList[T comparable] interface {
	Size() uint64
	Empty() bool
	ValueAt(index uint64) (T, error)
	ElementAt(index uint64) *ListElement[T]
	//PushFront(element *ListElement)
	//PopFront() *ListElement
	PushBack(element *ListElement[T])
	PopBack() *ListElement[T]
	//Front() *ListElement
	//Back() *ListElement
	//InsertAt(element *ListElement, index uint64) error
	//Erase(index uint64) error
	// ValueAtFromEnd(index uint64) int64
	// Reverse()
	//RemoveValue(value int64) error
	//Print()
}

// NewEmptyLinkedList Creates a new linkedlist
//func NewEmptyLinkedList[T comparable]() LinkedList[T] {
//	return &linkedList[T]{}
//}

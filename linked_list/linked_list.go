package linkedlist

// LinkedList Represnets the interface to access the linkedList elements
type LinkedList interface {
	Size() int64
	Empty() bool
	PushBack(element *ListElement)
	PushFront(element *ListElement)
}

// NewEmptyLinkedList Creates a new linkedlist
func NewEmptyLinkedList() LinkedList {
	return &linkedList{}
}

type linkedList struct {
	head *ListElement
	tail *ListElement
	size int64
}

func (list *linkedList) Size() int64 {
	return 0
}

func (list *linkedList) Empty() bool {
	return false
}

func (list *linkedList) PushBack(element *ListElement) {}

func (list *linkedList) PushFront(element *ListElement) {}

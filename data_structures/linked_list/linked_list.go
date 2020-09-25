package linkedlist

import "fmt"

// LinkedList Represents the interface to access the linkedList elements
type LinkedList interface {
	Size() uint64
	Empty() bool
	ValueAt(index uint64) (int64, error)
	PushBack(element *ListElement)
	PushFront(element *ListElement)
	Front() *ListElement
	Back() *ListElement
	Print()
}

// NewEmptyLinkedList Creates a new linkedlist
func NewEmptyLinkedList() LinkedList {
	return &linkedList{}
}

type linkedList struct {
	head *ListElement
	tail *ListElement
	size uint64
}

func (list *linkedList) Size() uint64 {
	return list.size
}

func (list *linkedList) Empty() bool {
	return list.size == 0
}

func (list *linkedList) ValueAt(index uint64) (int64, error) {
	if err := list.validIndex(index); err != nil {
		return 0, err
	}

	node := list.elementAt(index)

	return node.GetValue(), nil
}

func (list *linkedList) PushBack(element *ListElement) {
	if element == nil {
		return
	}

	if list.size == 0 {
		list.head = element
		list.tail = element
		list.size++
		return
	}

	list.tail.SetNext(element)
	list.tail = element
	list.size++
}

func (list *linkedList) PushFront(element *ListElement) {
	if element == nil {
		return
	}

	if list.size == 0 {
		list.head = element
		list.tail = element
		list.size++
		return
	}

	element.SetNext(list.head)
	list.head = element
	list.size++
}

func (list *linkedList) Front() *ListElement {
	return list.head
}

func (list *linkedList) Back() *ListElement {
	return list.tail
}

func (list *linkedList) Print() {
	node := list.head
	fmt.Print("[front] ")
	count := 0
	for node != nil {
		fmt.Printf("[%d] %d ->", count, node.GetValue())
		node = node.next
		count++
	}
	fmt.Println()
}

func (list *linkedList) validIndex(index uint64) error {
	// list index between: [0, list.size - 1]
	if index >= list.size {
		return fmt.Errorf("index out of bounds")
	}

	return nil
}

func (list *linkedList) elementAt(index uint64) *ListElement {
	// last element (tail)
	if index == list.size-1 {
		return list.tail
	}

	// scroll through linked list, including the head case
	node := list.head
	var position uint64 = 0
	for node != nil {
		if position == index {
			break
		}

		node = node.GetNext()
		position++
	}

	return node
}

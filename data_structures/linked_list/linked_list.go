package linkedlist

import "fmt"

// LinkedList Represents the interface to access the linkedList elements
type LinkedList interface {
	Size() int64
	Empty() bool
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
	size int64
}

func (list *linkedList) Size() int64 {
	return list.size
}

func (list *linkedList) Empty() bool {
	return list.size == 0
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

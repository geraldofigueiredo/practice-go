package linkedlist

import (
	"errors"
	"fmt"
)

// LinkedList Represents the interface to access the linkedList elements
type LinkedList interface {
	Size() uint64
	Empty() bool
	ValueAt(index uint64) (int64, error)
	ElementAt(index uint64) *ListElement
	PushFront(element *ListElement)
	PopFront() *ListElement
	PushBack(element *ListElement)
	PopBack() *ListElement
	Front() *ListElement
	Back() *ListElement
	InsertAt(element *ListElement, index uint64) error
	Erase(index uint64) error
	// ValueAtFromEnd(index uint64) int64
	// Reverse()
	RemoveValue(value int64) error
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

	node := list.ElementAt(index)

	return node.GetValue(), nil
}

func (list *linkedList) PushBack(element *ListElement) {
	if element == nil {
		return
	}

	if list.Empty() {
		list.head = element
		list.tail = element
		list.size++
		return
	}

	list.tail.SetNext(element)
	list.tail = element
	list.size++
}

func (list *linkedList) PopBack() *ListElement {
	if list.tail == nil {
		return nil
	}

	var element *ListElement = list.head
	var previous *ListElement = nil
	for element != list.tail {
		previous = element
		element = element.GetNext()
	}

	list.tail = previous
	list.size--

	return element
}

func (list *linkedList) PushFront(element *ListElement) {
	if element == nil {
		return
	}

	if list.Empty() {
		list.head = element
		list.tail = element
		list.size++
		return
	}

	element.SetNext(list.head)
	list.head = element
	list.size++
}

func (list *linkedList) PopFront() *ListElement {
	if list.head == nil {
		return nil
	}

	element := list.head
	list.head = list.head.GetNext()
	list.size--

	return element
}

func (list *linkedList) InsertAt(element *ListElement, index uint64) error {
	if element == nil {
		return errors.New("nil element")
	}

	if err := list.validIndex(index); err != nil {
		return errors.New("index out of bounds")
	}

	if list.indexAtHead(index) {
		list.PushFront(element)
	}

	if list.indexAtTail(index) {
		list.PushBack(element)
	}

	return nil
}

func (list *linkedList) Erase(index uint64) error {
	if list.head == nil {
		return fmt.Errorf("empty List")
	}

	var element *ListElement = list.head
	var previous *ListElement = nil
	var count uint64 = 0
	for count != index && count <= list.size {
		previous = element
		element = element.GetNext()
		count++
	}

	if element == nil {
		return fmt.Errorf("value not found")
	}

	previous = element.GetNext()
	list.size--
	return nil
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

func (list *linkedList) ElementAt(index uint64) *ListElement {
	// last element (tail)
	if index == list.size-1 {
		return list.tail
	}

	// scroll through linked list, including the head case
	node := list.head
	var position uint64 = 0
	for node != nil {
		if position == index {
			return node
		}

		node = node.GetNext()
		position++
	}

	return nil
}

func (list *linkedList) validIndex(index uint64) error {
	// list index between: [0, list.size - 1]
	if index >= list.size {
		return fmt.Errorf("index out of bounds")
	}

	return nil
}

func (list *linkedList) RemoveValue(value int64) error {
	if list.head == nil {
		return fmt.Errorf("empty List")
	}

	var element *ListElement = list.head
	var previous *ListElement = nil
	for element == nil || element.GetValue() != value {
		previous = element
		element = element.GetNext()
	}

	if element == nil {
		return fmt.Errorf("value not found")
	}

	previous = element.GetNext()
	list.size--
	return nil
}

func (list *linkedList) GetHead() *ListElement {
	return list.head
}

func (list *linkedList) indexAtHead(index uint64) bool {
	return index == 0
}

func (list *linkedList) indexAtTail(index uint64) bool {
	return index == list.size-1
}

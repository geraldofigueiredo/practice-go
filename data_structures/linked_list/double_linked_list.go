package linkedlist

import (
	"fmt"
)

type doubleLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size uint64
}

func NewDoubleLinkedList[T comparable]() DoubleLinkedList[T] {
	return &doubleLinkedList[T]{}
}

func (list *doubleLinkedList[_]) Size() uint64 {
	return list.size
}

func (list *doubleLinkedList[_]) Empty() bool {
	return list.size == 0
}

func (list *doubleLinkedList[T]) Head() *Node[T] {
	return list.head
}

func (list *doubleLinkedList[T]) Tail() *Node[T] {
	return list.tail
}

func (list *doubleLinkedList[T]) ValueAt(index uint64) (data T, err error) {
	if err := list.validIndex(index); err != nil {
		return data, err
	}

	node, err := list.ElementAt(index)
	if err != nil {
		var zeroValue T
		return zeroValue, err
	}

	return node.GetValue(), nil
}

// PushBack add a new element at the end
func (list *doubleLinkedList[T]) PushBack(value T) {
	element := NewNode(value, nil, nil)

	// empty list, the new element are the head and tail
	if list.Empty() {
		element.SetNext(nil)
		element.SetPrevious(nil)
		list.head = element
		list.tail = element
		list.size++
		return
	}

	// add the actual tail as the previous to the new element
	element.SetNext(nil)
	element.SetPrevious(list.tail)

	// update the tail with the current tail element
	list.tail.SetNext(element)
	list.tail = element
	list.size++
}

// remove the element at the end of the list
func (list *doubleLinkedList[T]) PopBack() *Node[T] {
	if list.Empty() {
		return nil
	}

	if list.size == 1 {
		element := list.head
		list.head = nil
		list.tail = nil
		list.size = 0
		return element
	}

	lastElement := list.tail

	list.tail = list.tail.previous
	list.size--

	return lastElement
}

func (list *doubleLinkedList[T]) PushFront(value T) {
	element := NewNode(value, nil, nil)

	if list.Empty() {
		list.head = element
		list.tail = element
		list.size++
		return
	}

	element.SetNext(list.head)
	element.SetPrevious(nil)
	list.head = element
	list.size++
}

func (list *doubleLinkedList[T]) PopFront() *Node[T] {
	if list.Empty() {
		return nil
	}

	if list.size == 1 {
		element := list.head
		list.head = nil
		list.tail = nil
		list.size = 0
		return element
	}

	element := list.head
	list.head = list.head.GetNext()
	list.head.SetPrevious(nil)
	list.size--

	return element
}

func (list *doubleLinkedList[T]) Print() {
	node := list.head
	fmt.Print("[front] ")
	count := 0
	for node != nil {
		fmt.Printf("[%d] %v ->", count, node.GetValue())
		node = node.next
		count++
	}
	fmt.Println()
}

func (list *doubleLinkedList[T]) ElementAt(index uint64) (*Node[T], error) {
	if err := list.validIndex(index); err != nil {
		return nil, err
	}

	// last element (tail)
	if index == list.size-1 {
		return list.tail, nil
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

	return node, nil
}

// Find should find the node searching by value and return the element and index
func (list *doubleLinkedList[T]) Find(value T) (*Node[T], uint64) {
	if list.Empty() {
		return nil, 0
	}

	// this validation cover the one node list case
	if list.tail.GetValue() == value {
		return list.tail, list.size - 1
	}

	element := list.head
	index := uint64(0)
	for element != list.tail {
		if element.GetValue() == value {
			return element, index
		}

		element = element.GetNext()
		index++
	}

	return nil, 0
}

func (list *doubleLinkedList[T]) RemoveAt(index uint64) (*Node[T], error) {
	if err := list.validIndex(index); err != nil {
		return nil, err
	}

	// remove head
	if index == uint64(0) {
		element := list.head
		list.head = list.head.next
		list.size--
		return element, nil
	}

	// remove tail
	if index == list.size-1 {
		element := list.tail
		list.tail = element.previous
		list.size--
		return element, nil
	}

	node := list.head
	position := uint64(0)
	for node != nil {
		node = node.GetNext()
		position++

		if position == index {
			break
		}
	}

	previous := node.previous
	next := node.next

	// remove from the middle of the list
	previous.next = next
	next.previous = previous

	return node, nil
}

func (list *doubleLinkedList[T]) validIndex(index uint64) error {
	// list index between: [0, list.size - 1]
	if index >= list.size {
		return fmt.Errorf("index out of bounds")
	}

	return nil
}

//
//// TODO: Fix remove value method
//func (list *doubleLinkedList) RemoveValue(value int64) error {
//	// if list.head == nil {
//	// 	return fmt.Errorf("empty List")
//	// }
//
//	// var element *Node = list.head
//	// var previous *Node = nil
//	// for element == nil || element.GetValue() != value {
//	// 	previous = element
//	// 	element = element.GetNext()
//	// }
//
//	// if element == nil {
//	// 	return fmt.Errorf("value not found")
//	// }
//
//	// previous = element.GetNext()
//	// list.size--
//	return nil
//}
//
//func (list *doubleLinkedList) GetHead() *Node {
//	return list.head
//}
//
//func (list *doubleLinkedList) indexAtHead(index uint64) bool {
//	return index == 0
//}
//
//func (list *doubleLinkedList) indexAtTail(index uint64) bool {
//	return index == list.size-1
//}

package linkedlist

//
//import "fmt"
//
//type linkedList[T comparable] struct {
//	head *ListElement[T]
//	tail *ListElement[T]
//	size uint64
//}
//
//func (list *linkedList[_]) Size() uint64 {
//	return list.size
//}
//
//func (list *linkedList[_]) Empty() bool {
//	return list.size == 0
//}
//
//func (list *linkedList[T]) ValueAt(index uint64) (data T, err error) {
//	if err := list.validIndex(index); err != nil {
//		return data, err
//	}
//
//	node := list.ElementAt(index)
//
//	return node.GetValue(), nil
//}
//
//func (list *linkedList[T]) PushBack(element *ListElement[T]) {
//	if element == nil {
//		return
//	}
//
//	if list.Empty() {
//		list.head = element
//		list.tail = element
//		list.size++
//		return
//	}
//
//	list.tail.SetNext(element)
//	list.tail = element
//	list.size++
//}
//
//func (list *linkedList[T]) PopBack() *ListElement[T] {
//	if list.tail == nil {
//		return nil
//	}
//
//	var element *ListElement[T] = list.head
//	var previous *ListElement[T] = nil
//	for element != list.tail {
//		previous = element
//		element = element.GetNext()
//	}
//
//	list.tail = previous
//	list.size--
//
//	return element
//}
//
////	func (list *linkedList) PushFront(element *ListElement) {
////		if element == nil {
////			return
////		}
////
////		if list.Empty() {
////			list.head = element
////			list.tail = element
////			list.size++
////			return
////		}
////
////		element.SetNext(list.head)
////		list.head = element
////		list.size++
////	}
////
////	func (list *linkedList) PopFront() *ListElement {
////		if list.head == nil {
////			return nil
////		}
////
////		element := list.head
////		list.head = list.head.GetNext()
////		list.size--
////
////		return element
////	}
////
////	func (list *linkedList) InsertAt(element *ListElement, index uint64) error {
////		if element == nil {
////			return errors.New("nil element")
////		}
////
////		if err := list.validIndex(index); err != nil {
////			return errors.New("index out of bounds")
////		}
////
////		if list.indexAtHead(index) {
////			list.PushFront(element)
////		}
////
////		if list.indexAtTail(index) {
////			list.PushBack(element)
////		}
////
////		return nil
////	}
////
//// // TODO: Fix erase method
////
////	func (list *linkedList) Erase(index uint64) error {
////		// if list.head == nil {
////		// 	return fmt.Errorf("empty List")
////		// }
////
////		// var element *ListElement = list.head
////		// var previous *ListElement = nil
////		// var count uint64 = 0
////		// for count != index && count <= list.size {
////		// 	previous = element
////		// 	element = element.GetNext()
////		// 	count++
////		// }
////
////		// if element == nil {
////		// 	return fmt.Errorf("value not found")
////		// }
////
////		// previous = element.GetNext()
////		// list.size--
////		return nil
////	}
////
////	func (list *linkedList) Front() *ListElement {
////		return list.head
////	}
////
////	func (list *linkedList) Back() *ListElement {
////		return list.tail
////	}
//func (list *linkedList[T]) Print() {
//	node := list.head
//	fmt.Print("[front] ")
//	count := 0
//	for node != nil {
//		fmt.Printf("[%d] %d ->", count, node.GetValue())
//		node = node.next
//		count++
//	}
//	fmt.Println()
//}
//
//func (list *linkedList[T]) ElementAt(index uint64) *ListElement[T] {
//	// last element (tail)
//	if index == list.size-1 {
//		return list.tail
//	}
//
//	// scroll through linked list, including the head case
//	node := list.head
//	var position uint64 = 0
//	for node != nil {
//		if position == index {
//			return node
//		}
//
//		node = node.GetNext()
//		position++
//	}
//
//	return nil
//}
//
//func (list *linkedList[T]) validIndex(index uint64) error {
//	// list index between: [0, list.size - 1]
//	if index >= list.size {
//		return fmt.Errorf("index out of bounds")
//	}
//
//	return nil
//}
//
////
////// TODO: Fix remove value method
////func (list *linkedList) RemoveValue(value int64) error {
////	// if list.head == nil {
////	// 	return fmt.Errorf("empty List")
////	// }
////
////	// var element *ListElement = list.head
////	// var previous *ListElement = nil
////	// for element == nil || element.GetValue() != value {
////	// 	previous = element
////	// 	element = element.GetNext()
////	// }
////
////	// if element == nil {
////	// 	return fmt.Errorf("value not found")
////	// }
////
////	// previous = element.GetNext()
////	// list.size--
////	return nil
////}
////
////func (list *linkedList) GetHead() *ListElement {
////	return list.head
////}
////
////func (list *linkedList) indexAtHead(index uint64) bool {
////	return index == 0
////}
////
////func (list *linkedList) indexAtTail(index uint64) bool {
////	return index == list.size-1
////}

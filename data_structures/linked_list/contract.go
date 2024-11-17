package linkedlist

// LinkedList Represents the interface to access the linkedList elements
type DoubleLinkedList[T comparable] interface {
	Size() uint64
	Empty() bool
	ValueAt(index uint64) (T, error)
	ElementAt(index uint64) (*Node[T], error)
	PushFront(data T)
	PopFront() *Node[T]
	PushBack(data T)
	PopBack() *Node[T]
	RemoveAt(index uint64) (*Node[T], error)
	Find(value T) (*Node[T], uint64)

	Head() *Node[T]
	Tail() *Node[T]

	Print()
}

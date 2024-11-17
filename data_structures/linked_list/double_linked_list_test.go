package linkedlist

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newEmptyDoubleLinkedList[T comparable]() *doubleLinkedList[T] {
	return &doubleLinkedList[T]{}
}

func TestNewDoubleLinkedList(t *testing.T) {
	got := NewDoubleLinkedList[int]()
	want := &doubleLinkedList[int]{}

	if equal := reflect.DeepEqual(got, want); !equal {
		t.Errorf("newEmptyDoubleLinkedList, got %v want %v", got, want)
	}
}

func TestPushBack(t *testing.T) {
	t.Run("push to empty list", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		list.PushBack(10)

		element, _ := list.ElementAt(0)
		assert.Equal(t, uint64(1), list.Size())
		assert.Equal(t, 10, element.GetValue())
		assert.Nil(t, element.GetNext())
		assert.Nil(t, element.GetPrevious())
	})

	t.Run("push an non empty list", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		headElement := NewNode(10, nil, nil)
		list.head = headElement
		list.tail = headElement
		list.size++

		list.PushBack(99)

		element, _ := list.ElementAt(1)
		assert.Equal(t, uint64(2), list.Size())
		assert.Equal(t, 99, element.GetValue())
		assert.Equal(t, list.head, headElement)
		assert.Equal(t, list.tail, element)
	})
}

func TestPopBack(t *testing.T) {
	t.Run("pop an empty list", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()

		element := list.PopBack()

		assert.Nil(t, element)
	})

	t.Run("pop a list with one single node, the list should be empty", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		element := NewNode(100, nil, nil)
		list.head = element
		list.tail = element
		list.size++

		removedElement := list.PopBack()

		assert.Equal(t, uint64(0), list.Size())
		assert.NotNil(t, removedElement)
		assert.Equal(t, element, removedElement)
		assert.True(t, list.Empty())
	})

	t.Run("pop a list with two nodes, the head and tail will be the same after pop", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		list.PushBack(500)
		list.PushBack(100)
		tail := list.tail

		removedElement := list.PopBack()

		assert.Equal(t, uint64(1), list.Size())
		assert.NotNil(t, removedElement)
		assert.Equal(t, tail, removedElement)
		assert.Equal(t, list.head, list.tail)
	})

	t.Run("pop a list with many nodes, should only change the tail", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		list.PushBack(500)
		list.PushBack(100)
		list.PushBack(999)
		list.PushBack(10)
		tail := list.tail

		removedElement := list.PopBack()

		assert.Equal(t, uint64(3), list.Size())
		assert.NotNil(t, removedElement)
		assert.Equal(t, tail, removedElement)
		assert.Equal(t, tail.previous, list.tail)
	})
}

func TestPushFront(t *testing.T) {
	t.Run("push front an empty list", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		list.PushFront(10)

		element, _ := list.ElementAt(0)
		assert.Equal(t, uint64(1), list.Size())
		assert.Equal(t, 10, element.GetValue())
		assert.Nil(t, element.next)
		assert.Nil(t, element.previous)
		assert.Equal(t, list.head, element)
		assert.Equal(t, list.tail, element)
	})

	t.Run("push front an non list, should change only the head", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		list.PushBack(10)
		list.PushFront(99)

		element, _ := list.ElementAt(0)
		assert.Equal(t, uint64(2), list.Size())
		assert.Equal(t, 99, element.GetValue())
		assert.NotNil(t, element.next)
		assert.Nil(t, element.previous)
		assert.Equal(t, 10, element.GetNext().GetValue())
	})
}

func TestPopFront(t *testing.T) {
	t.Run("pop front an empty list", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()

		element := list.PopFront()

		assert.Nil(t, element)
	})

	t.Run("pop front a list with one single node, should empty the list", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		list.PushBack(100)

		element := list.PopFront()

		assert.True(t, list.Empty())
		assert.Equal(t, 100, element.GetValue())
	})

	t.Run("pop front a list with many nodes, should change only the head of the list", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		list.PushBack(100)
		list.PushBack(200)
		list.PushBack(300)

		element := list.PopFront()

		assert.False(t, list.Empty())
		assert.Equal(t, 100, element.GetValue())
		assert.Equal(t, 200, list.head.GetValue())
	})
}

func TestSize(t *testing.T) {
	t.Run("empty size", func(*testing.T) {
		list := newEmptyDoubleLinkedList[int]()

		size := list.Size()

		assert.Equal(t, uint64(0), size)
	})

	t.Run("1 element", func(*testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		list.PushBack(100)

		assert.Equal(t, uint64(1), list.Size())

		list.PopBack()

		assert.Equal(t, uint64(0), list.Size())
	})
}

func TestFind(t *testing.T) {
	t.Run("find an empty list", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[string]()

		el, index := list.Find("some value")

		assert.Nil(t, el)
		assert.Zero(t, index)
	})

	t.Run("find a list with many nodes, and the searched value is the head of the list", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[string]()
		list.PushBack("some value")
		list.PushBack("awesome value")

		el, index := list.Find("some value")

		assert.NotNil(t, el)
		assert.Equal(t, list.head, el)
		assert.Equal(t, uint64(0), index)
	})

	t.Run("find a list with many nodes, and the searched value is the tail of the list", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[string]()
		list.PushBack("some value")
		list.PushBack("awesome value")

		el, index := list.Find("awesome value")

		assert.NotNil(t, el)
		assert.Equal(t, list.tail, el)
		assert.Equal(t, uint64(1), index)
	})

	t.Run("find a list with many nodes, and the searched value is in the middle of the list", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[string]()
		list.PushBack("some value 0")
		list.PushBack("some value 1")
		list.PushBack("awesome value")
		list.PushBack("some value 3")
		list.PushBack("some value 4")

		el, index := list.Find("awesome value")

		expected, _ := list.ElementAt(2)
		assert.NotNil(t, el)
		assert.Equal(t, expected, el)
		assert.Equal(t, uint64(2), index)
	})

	t.Run("find an element that not exist in the list", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[string]()
		list.PushBack("element 0")
		list.PushBack("element 1")
		list.PushBack("element 2")

		el, index := list.Find("element 99")

		assert.Nil(t, el)
		assert.Zero(t, index)
	})
}

func TestElementAt(t *testing.T) {
	t.Run("element invalid position, should return a nil", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[string]()
		list.PushBack("value 1")
		list.PushBack("other value")

		el, err := list.ElementAt(100)

		assert.Error(t, err)
		assert.Nil(t, el)
	})
}

func TestRemoveAt(t *testing.T) {
	t.Run("invalid position", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		list.PushBack(100)

		el, err := list.RemoveAt(99)

		assert.Nil(t, el)
		assert.Error(t, err)
	})

	t.Run("remove the head element", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		list.PushBack(100)
		head := list.head

		el, err := list.RemoveAt(0)

		assert.Equal(t, head, el)
		assert.NoError(t, err)
	})

	t.Run("remove the tail element", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		list.PushBack(100)
		list.PushBack(200)
		tail := list.tail

		el, err := list.RemoveAt(1)

		assert.Equal(t, tail, el)
		assert.NoError(t, err)
	})

	t.Run("remove in the middle", func(t *testing.T) {
		list := newEmptyDoubleLinkedList[int]()
		list.PushBack(100)
		list.PushBack(200)
		list.PushBack(300)
		list.PushBack(400)
		list.PushBack(500)

		want, _ := list.ElementAt(2)

		got, err := list.RemoveAt(2)

		assert.Equal(t, want, got)
		assert.NoError(t, err)
	})
}

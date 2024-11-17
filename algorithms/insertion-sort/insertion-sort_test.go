package insertionsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	t.Run("sort and empty list", func(t *testing.T) {
		var node *Node = nil

		got := InsertionSortList(node)

		assert.Nil(t, got)
	})

	t.Run("sort a single node list", func(t *testing.T) {
		var node *Node = &Node{
			Val:  1,
			Next: nil,
		}

		got := InsertionSortList(node)

		assert.Equal(t, node, got)
	})

	t.Run("sort a two nodes list", func(t *testing.T) {
		var head *Node = &Node{
			Val: 90,
			Next: &Node{
				Val:  2,
				Next: nil,
			},
		}

		want := &Node{
			Val: 2,
			Next: &Node{
				Val:  90,
				Next: nil,
			},
		}

		t.Log(want)

		got := InsertionSortList(head)

		assert.Equal(t, want, got)
	})
}

func TestSortedInsert(t *testing.T) {
	t.Run("insert an empty list", func(t *testing.T) {
		var n *Node = nil

		sorted := SortedInsert(&Node{Val: 10, Next: nil}, n)

		assert.Equal(t, 10, sorted.Val)
	})

	t.Run("insert a single node list", func(t *testing.T) {
		n := &Node{
			Val:  10,
			Next: nil,
		}

		sorted := SortedInsert(&Node{Val: 20, Next: nil}, n)

		assert.Equal(t, 10, sorted.Val)
	})

	t.Run("insert a multi node list", func(t *testing.T) {
		sorted := &Node{
			Val: 10,
			Next: &Node{
				Val: 20,
				Next: &Node{
					Val: 30,
					Next: &Node{
						Val:  100,
						Next: nil,
					},
				},
			},
		}

		want := &Node{
			Val: 10,
			Next: &Node{
				Val: 20,
				Next: &Node{
					Val: 30,
					Next: &Node{
						Val: 50,
						Next: &Node{
							Val:  100,
							Next: nil,
						},
					},
				},
			},
		}

		SortedInsert(&Node{Val: 50, Next: nil}, sorted)

		assert.Equal(t, want, sorted)
	})
}

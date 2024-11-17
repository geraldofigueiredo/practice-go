package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	t.Run("should put a new value in my LRU cache", func(t *testing.T) {
		lru := Constructor(10)
		lru.Put(10, 100)

		assert.Equal(t, 1, lru.size)
		assert.Equal(t, uint64(1), lru.usageList.Size())
		assert.Equal(t, 1, len(lru.cache))
	})

	t.Run("should put some valeus and retrieve the first value, should move the first value added to the head of the list", func(t *testing.T) {
		lru := Constructor(10)
		lru.Put(10, 100)
		lru.Put(20, 200)
		lru.Put(30, 300)
		// 10, 20, 30

		assert.Equal(t, 3, len(lru.cache))
		assert.Equal(t, 10, lru.usageList.Head().GetValue())
		assert.Equal(t, 30, lru.usageList.Tail().GetValue())

		lru.Get(10)
		// 20, 30, 10

		assert.Equal(t, 3, len(lru.cache))
		assert.Equal(t, 20, lru.usageList.Head().GetValue())
		assert.Equal(t, 10, lru.usageList.Tail().GetValue())
	})

	t.Run("should put some valeus and retrieve the first value, should move the first value added to the head of the list", func(t *testing.T) {
		lru := Constructor(3)
		lru.Put(10, 100)
		lru.Put(20, 200)
		lru.Put(30, 300)

		// 10, 20, 30 -> full cache

		assert.Equal(t, 3, len(lru.cache))
		assert.Equal(t, 3, lru.size)
		assert.Equal(t, 10, lru.usageList.Head().GetValue())
		assert.Equal(t, 30, lru.usageList.Tail().GetValue())

		lru.Put(50, 500)
		// 20, 30, 50

		assert.Equal(t, 20, lru.usageList.Head().GetValue())
		assert.Equal(t, 50, lru.usageList.Tail().GetValue())
		assert.Equal(t, -1, lru.Get(10))
	})

	t.Run("test case 1", func(t *testing.T) {
		lru := Constructor(2)
		lru.Put(1, 1)
		lru.Put(2, 2)

		// 1, 2

		lru.Get(1)

		// 2, 1

		assert.Equal(t, 2, lru.usageList.Head().GetValue())
		assert.Equal(t, 1, lru.usageList.Tail().GetValue())

		lru.Put(3, 3)

		// 1, 3

		assert.Equal(t, 1, lru.usageList.Head().GetValue())
		assert.Equal(t, 3, lru.usageList.Tail().GetValue())

		assert.Equal(t, -1, lru.Get(2))

		lru.Put(4, 4)

		// 3, 4

		assert.Equal(t, 3, lru.usageList.Head().GetValue())
		assert.Equal(t, 4, lru.usageList.Tail().GetValue())

		assert.Equal(t, -1, lru.Get(1))
		assert.Equal(t, 3, lru.Get(3))
		assert.Equal(t, 4, lru.Get(4))
	})

	t.Run("test case 15", func(t *testing.T) {
		lru := Constructor(2)
		lru.Get(2)    // -1
		lru.Put(2, 6) // 2=6

		assert.Equal(t, 2, lru.usageList.Head().GetValue())
		assert.Equal(t, 2, lru.usageList.Tail().GetValue())

		lru.Get(1)    // -1
		lru.Put(1, 5) // 2=6, 1=5

		assert.Equal(t, 2, lru.usageList.Head().GetValue())
		assert.Equal(t, 1, lru.usageList.Tail().GetValue())

		lru.Put(1, 2) // 2=6, 1=2

		assert.Equal(t, 2, lru.usageList.Head().GetValue())
		assert.Equal(t, 1, lru.usageList.Tail().GetValue())

		assert.Equal(t, 2, lru.Get(1))
		assert.Equal(t, 6, lru.Get(2))

		// assert.Equal(t, -1, lru.Get(2))

		// lru.Put(4, 4)

		// // 3, 4

		// assert.Equal(t, 3, lru.usageList.Head().GetValue())
		// assert.Equal(t, 4, lru.usageList.Tail().GetValue())

		// assert.Equal(t, -1, lru.Get(1))
		// assert.Equal(t, 3, lru.Get(3))
		// assert.Equal(t, 4, lru.Get(4))
	})
}

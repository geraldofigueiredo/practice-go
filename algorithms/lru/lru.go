package lru

import linkedlist "github.com/geraldofigueiredo/practice-go/data_structures/linked_list"

type LRUCache struct {
	capacity  int
	size      int
	cache     map[int]int
	usageList linkedlist.DoubleLinkedList[int]
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		size:      0,
		cache:     make(map[int]int),
		usageList: linkedlist.NewDoubleLinkedList[int](),
	}
}

func (l *LRUCache) Get(key int) int {
	value, ok := l.cache[key]
	if !ok {
		return -1
	}

	// search the element at the list
	node, position := l.usageList.Find(key)
	// remove the element and move to tail
	l.usageList.RemoveAt(position)
	l.usageList.PushBack(node.GetValue())

	return value
}

func (l *LRUCache) Put(key int, value int) {
	// verify if the element exists
	_, ok := l.cache[key]
	if ok {
		// only update the value at the cache
		l.cache[key] = value
		_, position := l.usageList.Find(key)
		l.usageList.RemoveAt(position)
		l.usageList.PushBack(key)
		return
	}

	// add a new position to the cache
	if l.size == l.capacity {
		l.evict()
	}

	// create the new key and add the usagelist
	l.cache[key] = value
	l.usageList.PushBack(key)
	l.size++
}

func (l *LRUCache) evict() {
	// remove the head of the list
	element, _ := l.usageList.RemoveAt(0)

	// clean the key in my map
	delete(l.cache, element.GetValue())

	// decrese my list size
	l.size--
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

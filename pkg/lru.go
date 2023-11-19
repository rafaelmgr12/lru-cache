package lru

import "container/list"

type LRUCache struct {
	capacity int
	items    map[int]*list.Element
	queue    *list.List
}

type cacheItem struct {
	key   int
	value interface{}
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}
func (c *LRUCache) Get(key int) interface{} {
	if elem, found := c.items[key]; found {
		c.queue.MoveToFront(elem)
		return elem.Value.(*cacheItem).value
	}
	return -1
}

func (c *LRUCache) Set(key int, value interface{}) {
	if elem, found := c.items[key]; found {
		c.queue.MoveToFront(elem)
		elem.Value.(*cacheItem).value = value
		return
	}

	if c.queue.Len() == c.capacity {
		oldest := c.queue.Back()
		delete(c.items, oldest.Value.(*cacheItem).key)
		c.queue.Remove(oldest)
	}
	elem := c.queue.PushFront(&cacheItem{key, value})
	c.items[key] = elem
}

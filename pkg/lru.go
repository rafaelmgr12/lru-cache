package lru

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	capacity int
	items    map[int]*list.Element
	queue    *list.List
	mu       sync.RWMutex // Usando RWMutex para leitura otimizada
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
	c.mu.RLock() // Lock para leitura
	defer c.mu.RUnlock()
	if elem, found := c.items[key]; found {
		c.queue.MoveToFront(elem)
		return elem.Value.(*cacheItem).value
	}
	return -1
}

func (c *LRUCache) Set(key int, value interface{}) {
	newItem := &cacheItem{key, value} // Criar o item fora do bloqueio

	c.mu.Lock() // Lock para escrita
	defer c.mu.Unlock()
	if elem, found := c.items[key]; found {
		c.queue.MoveToFront(elem)
		elem.Value = newItem
		return
	}

	if c.queue.Len() == c.capacity {
		oldest := c.queue.Back()
		delete(c.items, oldest.Value.(*cacheItem).key)
		c.queue.Remove(oldest)
	}
	elem := c.queue.PushFront(newItem)
	c.items[key] = elem
}

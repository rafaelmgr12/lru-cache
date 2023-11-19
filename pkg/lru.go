package lru

import (
	"container/list"
	"encoding/json"
	"io/ioutil"
	"sync"
)

type LRUCache struct {
	capacity int
	items    map[int]*list.Element
	queue    *list.List
	mu       sync.RWMutex
}

type cacheItem struct {
	Key   int         `json:"key"`
	Value interface{} `json:"value"`
}

type cacheState struct {
	Items []cacheItem `json:"items"`
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}

func (c *LRUCache) Get(key int) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if elem, found := c.items[key]; found {
		c.queue.MoveToFront(elem)
		return elem.Value.(*cacheItem).Value
	}
	return -1
}

func (c *LRUCache) Set(key int, value interface{}) {
	newItem := &cacheItem{key, value}

	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, found := c.items[key]; found {
		c.queue.MoveToFront(elem)
		elem.Value = newItem
		return
	}

	if c.queue.Len() == c.capacity {
		oldest := c.queue.Back()
		delete(c.items, oldest.Value.(*cacheItem).Key)
		c.queue.Remove(oldest)
	}
	elem := c.queue.PushFront(newItem)
	c.items[key] = elem
}

func (c *LRUCache) SaveToFile(filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var state cacheState
	for elem := c.queue.Front(); elem != nil; elem = elem.Next() {
		item := elem.Value.(*cacheItem)
		state.Items = append(state.Items, *item)
	}

	data, err := json.Marshal(state)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}
func (c *LRUCache) LoadFromFile(filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var state cacheState
	if err := json.Unmarshal(data, &state); err != nil {
		return err
	}

	c.items = make(map[int]*list.Element)
	c.queue.Init()
	for _, item := range state.Items {
		elem := c.queue.PushFront(&cacheItem{item.Key, item.Value})
		c.items[item.Key] = elem
	}

	return nil
}

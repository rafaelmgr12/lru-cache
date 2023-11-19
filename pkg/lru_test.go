package lru_test

import (
	"testing"

	lru "github.com/rafaelmgr12/lru-cache/pkg"
	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	cache := lru.NewLRUCache(2)

	cache.Set(1, "item1")
	cache.Set(2, "item2")

	t.Run("get existing item", func(t *testing.T) {
		item := cache.Get(1)
		assert.Equal(t, "item1", item)
	})

	t.Run("get non-existing item", func(t *testing.T) {
		item := cache.Get(3)
		assert.Equal(t, -1, item)
	})

	t.Run("get item after it was moved to front", func(t *testing.T) {
		cache.Get(1)
		cache.Set(3, "item3")

		item := cache.Get(2)
		assert.Equal(t, -1, item)

		item = cache.Get(1)
		assert.Equal(t, "item1", item)
	})
	t.Run("update existing item", func(t *testing.T) {
		cache.Set(1, "updatedItem1")
		item := cache.Get(1)
		assert.Equal(t, "updatedItem1", item, "Expected item to be updated")
	})

	// Test LRU property with multiple accesses
	t.Run("lru property with multiple accesses", func(t *testing.T) {
		cache.Set(2, "item2") // Reset item 2
		cache.Get(1)          // Access item 1
		cache.Set(3, "item3") // Add new item, should evict item 2

		item := cache.Get(2)
		assert.Equal(t, -1, item, "Expected item 2 to be evicted")

		item = cache.Get(1)
		assert.Equal(t, "updatedItem1", item, "Expected item 1 to remain")
	})

	// Test capacity constraint
	t.Run("capacity constraint", func(t *testing.T) {
		cache.Set(4, "item4") // This should evict item 3
		cache.Set(5, "item5") // Add another item

		item := cache.Get(3)
		assert.Equal(t, -1, item, "Expected item 3 to be evicted due to capacity constraint")

		item = cache.Get(4)
		assert.Equal(t, "item4", item, "Expected item 4 to be in cache")
	})

	// Test insertion of duplicates
	t.Run("insertion of duplicates", func(t *testing.T) {
		cache.Set(4, "newItem4") // Update existing item
		cache.Set(6, "item6")    // Add new item, should evict item 5

		item := cache.Get(5)
		assert.Equal(t, -1, item, "Expected item 5 to be evicted after inserting a new item")

		item = cache.Get(4)
		assert.Equal(t, "newItem4", item, "Expected updated item 4 to be in cache")
	})
}

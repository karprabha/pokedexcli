package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu      sync.RWMutex
	Entries map[string]CacheEntry
}

type CacheEntry struct {
	CreatedAt time.Time
	Data      []byte
}

func (c *Cache) Add(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entries[key] = CacheEntry{
		CreatedAt: time.Now(),
		Data:      data,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}

	return entry.Data, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.Entries {
			if time.Since(v.CreatedAt) > interval {
				delete(c.Entries, k)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Entries: make(map[string]CacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

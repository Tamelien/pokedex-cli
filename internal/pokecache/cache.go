package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{}
	newCache.entry = make(map[string]cacheEntry)
	newCache.interval = interval
	go newCache.reapLoop()
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.entry[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.entry {
			if time.Since(v.createdAt) > c.interval {
				delete(c.entry, k)
			}
		}
		c.mu.Unlock()
	}
}

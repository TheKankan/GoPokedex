package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cacheEntries: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	vcopy := append([]byte(nil), val...)

	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       vcopy,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.cacheEntries[key]
	if !exists {
		return nil, false
	} else {
		vcopy := append([]byte(nil), entry.val...)
		return vcopy, true
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()

		now := time.Now()
		for key, entry := range c.cacheEntries {
			if now.Sub(entry.createdAt) > interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}

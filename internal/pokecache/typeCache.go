package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	interval time.Duration
	entry    map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

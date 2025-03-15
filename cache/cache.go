package cache

import (
	"log"
	"sync"
	"time"
)

type CacheObject struct {
	StatusCode   int
	Headers      string
	ResponseBody string
	Created      time.Time
	ExpireAt     time.Time
}

type Cache struct {
	data map[string]*CacheObject
	mu   sync.RWMutex
}

// NewCache crea una instancia de caché

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]*CacheObject),
	}
}

func (c *Cache) Store(url *string, value *CacheObject) bool {
	if value == nil {
		return false
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[*url] = value
	return true
}

func (c *Cache) ClearAll() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]*CacheObject)
	log.Print("Clearing all cache objects...")
}

func (c *Cache) ClearRegister(url *string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, *url)
	return true
}

func (c *Cache) Get(url string) (*CacheObject, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cacheObject, found := c.data[url]
	if !found || time.Now().After(cacheObject.ExpireAt) {
		return nil, false // No encontrado o expirado
	}

	return cacheObject, true
}

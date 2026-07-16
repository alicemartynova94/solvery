package internal

import (
	"sync"
	"time"
)

type CacheTtl struct {
	mu *sync.RWMutex
	m  map[string]Element
}

type Element struct {
	value     interface{}
	expiredAt time.Time
}

func NewCache() *CacheTtl {
	return &CacheTtl{
		m: make(map[string]Element),
	}
}

func (c *CacheTtl) Set(key string, val interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = Element{
		value:     val,
		expiredAt: time.Now().Add(ttl),
	}
}

func (c *CacheTtl) Get(key string) (res interface{}, exist bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.m[key]

	if !ok || time.Now().After(v.expiredAt) {
		return nil, false
	}

	return v.value, true
}

func (c *CacheTtl) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.m {
		if time.Now().After(v.expiredAt) {
			delete(c.m, k)
		}
	}
}

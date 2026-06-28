package internal

import (
	"sync"
	"time"
)

type CacheTtl struct {
	mu   sync.RWMutex
	data map[interface{}]Type
}

type Type struct {
	val   interface{}
	expAt time.Time
}

func NewCache() *CacheTtl {
	return &CacheTtl{
		data: make(map[interface{}]Type),
	}
}

func (c *CacheTtl) Set(key, val interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = Type{
		val:   val,
		expAt: time.Now().Add(ttl),
	}
}

// при протухшем элементе удалять?
func (c *CacheTtl) Get(key interface{}) (res interface{}, exist bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[key]

	if !ok || time.Now().After(v.expAt) {
		return nil, false
	}

	return v.val, true
}

// очистить все или только протухшие?
func (c *CacheTtl) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.data {
		if time.Now().After(v.expAt) {
			delete(c.data, k)
		}
	}
}

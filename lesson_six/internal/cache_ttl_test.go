package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestCacheTtl_Set(t *testing.T) {
	cacheTtl := NewCache()
	cacheTtl.Set("one", "element", 5*time.Minute)
	v, ok := cacheTtl.Get("one")

	assert.Equal(t, "element", v)
	assert.Equal(t, true, ok)
}

func TestCacheTtl_GetExpired(t *testing.T) {
	cacheTtl := NewCache()
	cacheTtl.Set("one", "element", 5*time.Millisecond)
	time.Sleep(10 * time.Millisecond)

	v, ok := cacheTtl.Get("one")

	assert.Nil(t, v)
	assert.False(t, ok)
}

func TestCacheTtl_GetWrongKey(t *testing.T) {
	cacheTtl := NewCache()
	cacheTtl.Set("one", "element", 5*time.Millisecond)

	v, ok := cacheTtl.Get("two")

	assert.Nil(t, v)
	assert.False(t, ok)
}

func TestCacheTtl_Clear(t *testing.T) {
	cacheTtl := NewCache()
	cacheTtl.Set("expired", "element", 5*time.Millisecond)
	cacheTtl.Set("not expired", "element", 5*time.Minute)
	time.Sleep(15 * time.Millisecond)

	cacheTtl.Clear()
	v1, ok1 := cacheTtl.Get("expired")
	v2, ok2 := cacheTtl.Get("not expired")

	assert.Nil(t, v1)
	assert.False(t, ok1)

	assert.Equal(t, "element", v2)
	assert.True(t, ok2)
}

func TestCacheTtl_GetConcurrent(t *testing.T) {
	cacheTtl := NewCache()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i)
			cacheTtl.Set(key, i, 5*time.Minute)
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i)
			cacheTtl.Get(key)
		}(i)
	}
}

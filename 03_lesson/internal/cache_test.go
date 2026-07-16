package lesson_three

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache_Get(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		wantVal string
		wantOk  bool
	}{
		{"key not found", 90, "", false},
		{"key found", 3, "blue", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewCache[int, string](10)
			cache.Put(1, "red")
			cache.Put(2, "green")
			cache.Put(3, "blue")
			cache.Put(4, "yellow")

			val, ok := cache.Get(tt.input)
			assert.Equal(t, tt.wantVal, val)
			assert.Equal(t, tt.wantOk, ok)
		})
	}
}

func TestCache_Put_IsEmpty(t *testing.T) {
	cache := NewCache[int, string](10)

	cache.Put(1, "red")

	assert.Equal(t, 1, cache.head.key)
	assert.Equal(t, "red", cache.tail.value)
}

func TestCache_Put_ExistingKey(t *testing.T) {
	cache := NewCache[int, string](10)
	cache.Put(1, "red")
	cache.Put(2, "red")
	cache.Put(3, "red")
	cache.Put(1, "blue")

	assert.Equal(t, "blue", cache.head.value)
}

func TestCache_Put_InCapacity(t *testing.T) {
	cache := NewCache[int, string](10)
	cache.Put(1, "yellow")
	cache.Put(2, "red")
	cache.Put(3, "green")
	cache.Put(4, "blue")

	assert.Equal(t, "blue", cache.head.value)
	assert.Equal(t, "yellow", cache.tail.value)
}

func TestCache_Put_ExceedCapacity(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		wantHead string
		wantTail string
	}{
		{"capacity: 1", 1, "purple", "purple"},
		{"capacity: 5", 5, "purple", "yellow"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewCache[int, string](tt.capacity)
			cache.Put(1, "red")
			cache.Put(2, "yellow")
			cache.Put(3, "white")
			cache.Put(4, "black")
			cache.Put(5, "green")
			cache.Put(6, "purple")

			assert.Equal(t, tt.wantHead, cache.head.value)
			assert.Equal(t, tt.wantTail, cache.tail.value)
		})
	}
}

func TestCache_Put_CheckOrder(t *testing.T) {
	cache := NewCache[int, string](10)
	cache.Put(3, "white")
	cache.Put(4, "black")
	cache.Put(5, "green")
	cache.Put(6, "purple")

	actual := make([]string, 0, 6)
	for node := cache.head; node != nil; node = node.next {
		actual = append(actual, node.value)
	}

	assert.Equal(t, []string{"purple", "green", "black", "white"}, actual)
}

func TestCache_Size(t *testing.T) {
	cache := NewCache[int, string](10)
	cache.Put(1, "red")
	cache.Put(2, "yellow")
	cache.Put(3, "white")
	cache.Put(4, "black")
	cache.Put(5, "green")
	cache.Put(6, "purple")

	assert.Equal(t, 6, cache.Size())
}

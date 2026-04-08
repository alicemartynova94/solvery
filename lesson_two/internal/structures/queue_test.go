package lesson_two

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func newQueue[T comparable](values ...T) Queue[T] {
	q := Queue[T]{}
	for _, value := range values {
		q.Push(value)
	}
	return q
}

func TestQueue_Push(t *testing.T) {
	tests := []struct {
		name         string
		values       []string
		expected     []string
		expectedSize int
	}{
		{"append to empty queue", []string{"first"}, []string{"first"}, 1},
		{"several appends", []string{"first", "second", "third", "forth"}, []string{"first", "second", "third", "forth"}, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := newQueue(test.values...)
			assert.Equal(t, test.expected, q.GetValues())
			assert.Equal(t, test.expectedSize, q.Size())

		})
	}
}

func TestQueue_Pop(t *testing.T) {
	tests := []struct {
		name        string
		values      []string
		expectedVal string
		want        bool
		size        int
	}{
		{"pop from empty list", []string{}, "", false, 0},
		{"pop from queue with one el", []string{"first"}, "first", true, 0},
		{"pop from queue with several el", []string{"forth", "third", "second", "first"}, "forth", true, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := newQueue(test.values...)
			val, ok := q.Pop()
			assert.Equal(t, test.expectedVal, val)
			assert.Equal(t, test.want, ok)

		})
	}
}

func TestQueue_Clear(t *testing.T) {
	tests := []struct {
		name         string
		values       []int
		expected     []int
		expectedSize int
	}{
		{"clear from empty queue", []int{}, []int{}, 0},
		{"clear queue with values", []int{90, 5, 7, 90, 20, 6, 90}, []int{}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := newQueue(test.values...)
			q.Clear()
			assert.Equal(t, test.expected, q.GetValues())
			assert.Equal(t, test.expectedSize, q.ll.size)

		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	q := newQueue([]int{}...)
	assert.True(t, q.IsEmpty())
}

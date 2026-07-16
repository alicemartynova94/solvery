package structures_test

import (
	"github.com/stretchr/testify/assert"
	"solvery/06_lesson/internal/structures"
	"sync"
	"testing"
)

func newStack[T comparable](values ...T) *structures.Stack[T] {
	stack := structures.NewStack[T]()
	for _, v := range values {
		stack.Push(v)
	}
	return stack
}

func TestStack_Push(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
		size     int
	}{
		{"push to empty stack", []int{1}, []int{1}, 1},
		{"push several el", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := newStack(test.values...)
			assert.Equal(t, test.expected, s.elements)
			assert.Equal(t, test.size, s.Size())

		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
		size     int
	}{
		{"pop from empty stack", []int{}, nil, 0},
		{"pop several el", []int{1, 2, 3, 4}, []int{1, 2, 3}, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := newStack(test.values...)
			s.Pop()
			assert.Equal(t, test.expected, s.elements)
			assert.Equal(t, test.size, s.Size())
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected int
		want     bool
		size     int
	}{
		{"peek empty stack", []int{}, 0, false, 0},
		{"peek stack with el", []int{1, 2, 3, 4}, 4, true, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := newStack(test.values...)
			val, ok := s.Peek()
			assert.Equal(t, test.expected, val)
			assert.Equal(t, test.want, ok)
			assert.Equal(t, test.size, s.Size())
		})
	}
}

func TestStack_Clear(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
		size     int
	}{
		{"clear from empty stack", []int{}, nil, 0},
		{"clear stack with values", []int{1, 2, 3, 4, 5}, []int{}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := newStack(test.values...)
			s.Clear()
			assert.Equal(t, test.expected, s.elements)
			assert.Equal(t, test.size, s.Size())
		})
	}
}

func TestStack_Concurrency(t *testing.T) {
	stack := newStack[int]()
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			stack.Push(n)
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stack.Pop()
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stack.Peek()
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stack.Clear()
		}()
	}

	wg.Wait()
}

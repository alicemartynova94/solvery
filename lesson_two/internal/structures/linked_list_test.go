package lesson_two

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func newList[T comparable](values ...T) LinkedList[T] {
	ll := LinkedList[T]{}
	for _, value := range values {
		ll.Append(value)
	}
	return ll
}

func TestLinkedList_Append(t *testing.T) {
	tests := []struct {
		name         string
		values       []string
		expected     []string
		expectedSize int
	}{
		{"append to empty list", []string{}, []string{}, 0},
		{"several appends", []string{"first", "second", "third", "forth"}, []string{"first", "second", "third", "forth"}, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := newList(test.values...)
			assert.Equal(t, test.expected, list.GetValues())
			assert.Equal(t, test.expectedSize, list.size)

		})
	}
}

func TestLinkedList_PrependToEmptyList(t *testing.T) {
	list := LinkedList[string]{}
	list.Prepend("first")

	assert.Equal(t, "first", list.head.value)
	assert.Equal(t, "first", list.tail.value)
	assert.Equal(t, 1, list.size)
}

func TestLinkedList_PrependToList(t *testing.T) {
	list := LinkedList[string]{}
	list.Prepend("first")
	list.Prepend("second")
	list.Prepend("third")
	list.Prepend("forth")

	assert.Equal(t, "forth", list.head.value)
	assert.Equal(t, "first", list.tail.value)
	assert.Equal(t, []string{"forth", "third", "second", "first"}, list.GetValues())
	assert.Equal(t, 4, list.size)
}

func TestLinkedList_RemoveTail(t *testing.T) {
	tests := []struct {
		name         string
		values       []string
		expected     []string
		expectedSize int
	}{
		{"remove from empty list", []string{}, []string{}, 0},
		{"remove tail from list with one el", []string{"first"}, []string{}, 0},
		{"remove tail from list with several el", []string{"forth", "third", "second", "first"}, []string{"forth", "third", "second"}, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := newList(test.values...)
			list.RemoveTail()
			assert.Equal(t, test.expected, list.GetValues())
			assert.Equal(t, test.expectedSize, list.size)

		})
	}
}

func TestLinkedList_RemoveFront(t *testing.T) {
	tests := []struct {
		name         string
		values       []string
		expected     []string
		expectedSize int
	}{
		{"remove from empty list", []string{}, []string{}, 0},
		{"remove front from list with one el", []string{"first"}, []string{}, 0},
		{"remove front from list with several el", []string{"forth", "third", "second", "first"}, []string{"third", "second", "first"}, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := newList(test.values...)
			list.RemoveFront()
			assert.Equal(t, test.expected, list.GetValues())
			assert.Equal(t, test.expectedSize, list.size)

		})
	}
}

func TestLinkedList_FindVal(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected bool
	}{
		{"find val in empty list", []string{}, false},
		{"find val in list with several el", []string{"forth", "third", "second", "first"}, true},
		{"no such el in list", []string{"forth", "third", "second", "fifth"}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := newList(test.values...)
			ok := list.FindVal("first")
			assert.Equal(t, test.expected, ok)
		})
	}
}

func TestLinkedList_RemoveAll(t *testing.T) {
	tests := []struct {
		name         string
		values       []int
		expected     []int
		expectedSize int
	}{
		{"remove all from empty list", []int{}, []int{}, 0},
		{"remove all till empty", []int{90, 90}, []int{}, 0},
		{"remove one value from list", []int{90, 5, 7}, []int{5, 7}, 2},
		{"remove several values from list", []int{90, 5, 7, 90, 20, 6, 90}, []int{5, 7, 20, 6}, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := newList(test.values...)
			list.RemoveAll(90)
			assert.Equal(t, test.expected, list.GetValues())
			assert.Equal(t, test.expectedSize, list.size)

		})
	}
}

func TestLinkedList_Clear(t *testing.T) {
	tests := []struct {
		name         string
		values       []int
		expected     []int
		expectedSize int
	}{
		{"clear from empty list", []int{}, []int{}, 0},
		{"clear list with values", []int{90, 5, 7, 90, 20, 6, 90}, []int{}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := newList(test.values...)
			list.Clear()
			assert.Equal(t, test.expected, list.GetValues())
			assert.Equal(t, test.expectedSize, list.size)

		})
	}
}

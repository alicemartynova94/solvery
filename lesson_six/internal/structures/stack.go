package structures

import "sync"

type Stack[T comparable] struct {
	mu       *sync.RWMutex
	elements []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		mu: &sync.RWMutex{},
	}
}

func (s *Stack[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.elements)
}

func (s *Stack[T]) Push(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.elements = append(s.elements, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var zeroVal T
	if len(s.elements) == 0 {
		return zeroVal, false
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element, true
}

func (s *Stack[T]) Peek() (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var zeroVal T
	if len(s.elements) == 0 {
		return zeroVal, false
	}

	result := s.elements[len(s.elements)-1]

	return result, true
}

func (s *Stack[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.elements = s.elements[:0]
}

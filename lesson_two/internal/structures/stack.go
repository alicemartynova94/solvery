package lesson_two

type Stack[T comparable] struct {
	elements []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) Push(v T) {
	s.elements = append(s.elements, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zeroVal T
	if s.IsEmpty() {
		return zeroVal, false
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zeroVal T
	if s.IsEmpty() {
		return zeroVal, false
	}

	result := s.elements[len(s.elements)-1]

	return result, true
}

func (s *Stack[T]) Clear() {
	s.elements = s.elements[:0]
}

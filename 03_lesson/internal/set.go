package lesson_three

type Set[T comparable] struct {
	values map[T]struct{}
}

func (s *Set[T]) Union(set *Set[T]) {
	for k, _ := range set.values {
		s.values[k] = struct{}{}
	}
}

func (s *Set[T]) Intersection(set *Set[T]) *Set[T] {
	intersection := make(map[T]struct{})
	small := s.values
	big := set.values
	if len(small) > len(big) {
		big = s.values
		small = set.values
	}

	for k, _ := range small {
		if _, ok := big[k]; ok {
			intersection[k] = struct{}{}
		}
	}
	return &Set[T]{intersection}
}

func (s *Set[T]) Difference(set *Set[T]) *Set[T] {
	difference := make(map[T]struct{})
	for k, _ := range s.values {
		if _, ok := set.values[k]; !ok {
			difference[k] = struct{}{}
		}
	}

	return &Set[T]{difference}
}

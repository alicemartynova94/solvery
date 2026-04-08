package lesson_two

type Queue[T comparable] struct {
	ll LinkedList[T]
}

func (q *Queue[T]) IsEmpty() bool {
	return q.ll.head == nil
}

func (q *Queue[T]) Size() int {
	return q.ll.size
}

func (q *Queue[T]) GetValues() []T {
	return q.ll.GetValues()
}

func (q *Queue[T]) Push(v T) {
	q.ll.Append(v)
}

func (q *Queue[T]) Pop() (T, bool) {
	return q.ll.RemoveFront()
}

func (q *Queue[T]) Clear() {
	q.ll.Clear()
}

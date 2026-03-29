package lesson_two

type Queue[T comparable] struct {
	head *QueueEl[T]
	tail *QueueEl[T]
	size int
}

type QueueEl[T comparable] struct {
	value T
	next  *QueueEl[T]
	prev  *QueueEl[T]
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) GetValues() []T {
	var values []T
	for i := q.head; i != nil; i = i.next {
		values = append(values, i.value)
	}
	return values
}

func (q *Queue[T]) Push(v T) {
	e := &QueueEl[T]{value: v}
	if q.IsEmpty() {
		q.tail, q.head = e, e
	} else {
		q.tail.next = e
		e.prev = q.tail
		q.tail = e
	}
	q.size++
}

func (q *Queue[T]) Pop() (T, bool) {
	var zeroVal T
	if q.IsEmpty() {
		return zeroVal, false
	}

	result := q.head.value
	if q.size == 1 {
		q.head, q.tail = nil, nil
		q.size = 0
		return result, true
	}

	oldHead := q.head
	q.head = oldHead.next
	q.head.prev = nil
	oldHead.next = nil
	q.size--

	return result, true
}

func (q *Queue[T]) Clear() {
	for e := q.head; e != nil; {
		next := e.next
		e.prev, e.next = nil, nil
		e = next
	}

	q.head, q.tail = nil, nil
	q.size = 0
}

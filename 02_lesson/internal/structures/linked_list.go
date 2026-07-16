package lesson_two

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

type Node[T comparable] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedList[T]) GetValues() []T {
	values := []T{}
	for i := l.head; i != nil; i = i.next {
		values = append(values, i.value)
	}
	return values
}

func (l *LinkedList[T]) Append(v T) {
	e := &Node[T]{value: v}
	if l.IsEmpty() {
		l.tail, l.head = e, e
	} else {
		l.tail.next = e
		e.prev = l.tail
		l.tail = e
	}
	l.size++
}

func (l *LinkedList[T]) Prepend(v T) {
	e := &Node[T]{value: v}
	if l.IsEmpty() {
		l.tail, l.head = e, e
	} else {
		l.head.prev = e
		e.next = l.head
		l.head = e
	}
	l.size++
}

func (l *LinkedList[T]) RemoveTail() (T, bool) {
	var zeroVal T
	if l.IsEmpty() {
		return zeroVal, false
	}

	result := l.tail.value
	if l.size == 1 {
		l.head, l.tail = nil, nil
		l.size = 0
		return result, true
	}

	oldTail := l.tail
	l.tail = oldTail.prev
	l.tail.next = nil
	oldTail.next = nil
	l.size--

	return result, true
}

func (l *LinkedList[T]) RemoveFront() (T, bool) {
	var zeroVal T
	if l.IsEmpty() {
		return zeroVal, false
	}

	result := l.head.value
	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
		return result, true
	}

	oldHead := l.head
	l.head = oldHead.next
	l.head.prev = nil
	oldHead.next = nil
	l.size--

	return result, true
}

func (l *LinkedList[T]) FindVal(v T) bool {
	if l.IsEmpty() {
		return false
	}

	var result bool
	for e := l.head; e != nil; e = e.next {
		if e.value == v {
			result = true
		}
	}

	return result
}

func (l *LinkedList[T]) RemoveAll(v T) {
	if l.IsEmpty() {
		return
	}

	for e := l.head; e != nil; {
		next := e.next
		if e.value == v {
			if e == l.head {
				_, _ = l.RemoveFront()
			} else if e == l.tail {
				_, _ = l.RemoveTail()
			} else {
				e.prev.next = e.next
				e.next.prev = e.prev
				l.size--
			}
		}
		e = next
	}
}

func (l *LinkedList[T]) Clear() {
	l.head, l.tail = nil, nil
	l.size = 0
}

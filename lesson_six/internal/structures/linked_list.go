package structures

import "sync"

type LinkedList[T comparable] struct {
	mu   *sync.RWMutex
	head *NodeL[T]
	tail *NodeL[T]
	size int
}

type NodeL[T comparable] struct {
	value T
	next  *NodeL[T]
	prev  *NodeL[T]
}

func (l *LinkedList[T]) Size() int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.size
}

func (l *LinkedList[T]) GetValues() []T {
	l.mu.RLock()
	defer l.mu.RUnlock()

	values := []T{}
	for i := l.head; i != nil; i = i.next {
		values = append(values, i.value)
	}
	return values
}

func (l *LinkedList[T]) Append(v T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	e := &NodeL[T]{value: v}
	if l.head == nil {
		l.tail, l.head = e, e
	} else {
		l.tail.next = e
		e.prev = l.tail
		l.tail = e
	}
	l.size++
}

func (l *LinkedList[T]) Prepend(v T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	e := &NodeL[T]{value: v}
	if l.head == nil {
		l.tail, l.head = e, e
	} else {
		l.head.prev = e
		e.next = l.head
		l.head = e
	}
	l.size++
}

func (l *LinkedList[T]) RemoveTail() (T, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.removeTailUnlocked()
}

func (l *LinkedList[T]) RemoveFront() (T, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.removeFrontUnlocked()
}

func (l *LinkedList[T]) FindVal(v T) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.head == nil {
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
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.head == nil {
		return
	}

	for e := l.head; e != nil; {
		next := e.next
		if e.value == v {
			if e == l.head {
				_, _ = l.removeFrontUnlocked()
			} else if e == l.tail {
				_, _ = l.removeTailUnlocked()
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
	l.mu.Lock()
	defer l.mu.Unlock()
	l.head, l.tail = nil, nil
	l.size = 0
}

func (l *LinkedList[T]) removeTailUnlocked() (T, bool) {
	var zeroVal T
	if l.head == nil {
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

func (l *LinkedList[T]) removeFrontUnlocked() (T, bool) {
	var zeroVal T
	if l.head == nil {
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

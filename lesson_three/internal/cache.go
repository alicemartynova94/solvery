package lesson_three

type Cache[K comparable, V any] struct {
	nodes    map[K]*Node[K, V]
	head     *Node[K, V]
	tail     *Node[K, V]
	capacity int
}

type Node[K comparable, V any] struct {
	key   K
	value V
	next  *Node[K, V]
	prev  *Node[K, V]
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity < 1 {
		panic("capacity must be greater than 0")
	}
	return &Cache[K, V]{
		nodes:    make(map[K]*Node[K, V]),
		capacity: capacity,
	}
}

func (c *Cache[K, V]) IsEmpty() bool {
	return c.head == nil
}

func (c *Cache[K, V]) Size() int {
	return len(c.nodes)
}

func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	var zeroValue V
	if _, ok = c.nodes[key]; !ok {
		return zeroValue, false
	}

	v, _ := c.nodes[key]
	c.pushFront(v)

	return v.value, true
}

func (c *Cache[K, V]) Put(key K, value V) {
	if v, ok := c.nodes[key]; ok {
		if len(c.nodes) > 1 {
			c.pushFront(v)
		}
		v.value = value
		return
	}

	node := &Node[K, V]{
		key:   key,
		value: value,
	}

	if c.IsEmpty() {
		c.nodes[key] = node
		c.head = node
		c.tail = node
		return
	}

	initialLength := len(c.nodes)
	c.nodes[key] = node
	c.head.prev = node
	node.next = c.head
	c.head = node

	if initialLength == c.capacity {
		old := c.tail
		c.tail = c.tail.prev
		c.tail.next = nil
		delete(c.nodes, old.key)
	}
}

func (c *Cache[K, V]) pushFront(v *Node[K, V]) {
	prev := v.prev
	next := v.next
	if v == c.tail {
		c.tail = prev
	}

	prev.next = next
	if next != nil {
		next.prev = prev
	}
	v.next = c.head
	c.head.prev = v
	v.prev = nil
	c.head = v
}

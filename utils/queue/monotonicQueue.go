package queue

type Node[E any] struct {
	val   E
	left  *Node[E]
	right *Node[E]
}

type MonotonicQueue[E any] struct {
	head *Node[E]
	tail *Node[E]
	size int
}

func NewMonotonicQueue[E any]() *MonotonicQueue[E] {
	return &MonotonicQueue[E]{
		size: 0,
	}
}

func (m *MonotonicQueue[E]) Push(e E) {
	n := Node[E]{
		val: e,
	}
	if m.head == nil {
		m.head, m.tail = &n, &n
		m.size = 1
	} else {
		m.tail.right, m.tail = &n, &n
		m.size++
	}
}

func (m *MonotonicQueue[E]) Pop() *E {
	if m.head == nil {
		return nil
	} else {
		val := m.head.val
		m.head = m.head.right
		if m.head != nil {
			m.head.left = nil
		}
		m.size--
		return &val
	}
}

func (m *MonotonicQueue[E]) Peek() *E {
	if m.head == nil {
		return nil
	} else {
		return &m.head.val
	}
}

func (m *MonotonicQueue[E]) PeekEnd() *E {
	if m.tail == nil {
		return nil
	} else {
		return &m.tail.val
	}
}

func (m *MonotonicQueue[E]) Len() int {
	return m.size
}

func (m *MonotonicQueue[E]) Max() {

}

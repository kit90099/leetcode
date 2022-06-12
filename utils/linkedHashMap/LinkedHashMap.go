package linkedhashmap

import "fmt"

type node struct {
	prev *node
	next *node
	val  *mapEntry
}

type doubleList struct {
	head *node
	tail *node
	size int
}

func (list *doubleList) len() int {
	return list.size
}

func (list *doubleList) addHead(node *node) {
	list.head, node.next = node, list.head
	list.size++
}

func (list *doubleList) removeTail(node *node) {
	tail := list.tail
	prev := list.tail.prev
	prev.next = nil
	tail.prev, tail.next = nil, nil
	list.tail = list.tail.prev
	list.size--
}

func (list *doubleList) remove(n *node) {
	n.prev.next, n.next.prev = n.next, n.prev
	n.prev = nil
	n.next = nil

	list.size--
}

type LinkedHashMap struct {
	keyMap map[int](*node)
	list   doubleList
}

func Default() LinkedHashMap {
	keyMap := make(map[int](*node), 0)
	list := doubleList{
		size: 0,
	}
	return LinkedHashMap{
		keyMap: keyMap,
		list:   list,
	}
}

type mapEntry struct {
	key   int
	value int
}

func (m *LinkedHashMap) Len() int {
	return m.list.len()
}

func (m *LinkedHashMap) Put(key int, value int) {
	entry := mapEntry{
		key:   key,
		value: value,
	}

	n := node{
		val: &entry,
	}

	m.list.addHead(&n)
	m.keyMap[key] = &n
}

func (m *LinkedHashMap) Get(key int) *int {
	node, exist := m.keyMap[key]
	if exist {
		return &node.val.value
	} else {
		return nil
	}
}

func (m *LinkedHashMap) Remove(key int) {
	node, exist := m.keyMap[key]
	if exist {
		m.list.remove(node)
		delete(m.keyMap, key)
	}
}

func (m *LinkedHashMap) ToString() string {
	str := ""

	ptr := m.list.head
	for ptr != nil {
		str += fmt.Sprintf("{key: %d, value: %d}, ", ptr.val.key, ptr.val.value)
		ptr = ptr.next
	}
	str += fmt.Sprintf("------------- size: %d", m.list.size)

	return str
}

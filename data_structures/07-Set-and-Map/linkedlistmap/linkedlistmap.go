package linkedlistmap

import "LearnGo/data_structures/utils"

type node struct {
	key   interface{}
	value interface{}
	next  *node
}

type LinkedListMap struct {
	dummyHead *node
	size      int
}

func NewLinkedListMap() *LinkedListMap {
	return &LinkedListMap{
		dummyHead: &node{},
		size:      0,
	}
}

func (m *LinkedListMap) GetSize() int {
	return m.size
}

func (m *LinkedListMap) IsEmpty() bool {
	return m.size == 0
}

func (m *LinkedListMap) Add(key interface{}, value interface{}) {
	if !m.Contains(key) {
		m.dummyHead.next = &node{
			key:   key,
			value: value,
			next:  m.dummyHead.next,
		}
		m.size++
	}
}

func (m *LinkedListMap) Set(key interface{}, value interface{}) {
	cur := m.dummyHead.next
	for cur != nil {
		if utils.Compare(cur.key, key) == 0 {
			cur.value = value
		}
		cur = cur.next
	}
}

func (m *LinkedListMap) Get(key interface{}) interface{} {
	cur := m.dummyHead.next
	for cur != nil {
		if utils.Compare(cur.key, key) == 0 {
			return cur.value
		}
		cur = cur.next
	}
	return nil
}

func (m *LinkedListMap) Remove(key interface{}) {
	prev := m.dummyHead
	for prev.next != nil {
		if utils.Compare(prev.next.key, key) == 0 {
			delNode := prev.next
			prev.next = delNode.next
			delNode.next = nil
			m.size--
			return
		}
		prev = prev.next
	}
}

func (m *LinkedListMap) Contains(key interface{}) bool {
	cur := m.dummyHead.next
	for cur != nil {
		if utils.Compare(cur.key, key) == 0 {
			return true
		}
		cur = cur.next
	}
	return false
}

package bstmap

import "LearnGo/data_structures/utils"

type node struct {
	key   interface{}
	value interface{}
	left  *node
	right *node
}

type BSTMap struct {
	root *node
	size int
}

func NewBSTMap() *BSTMap {
	return &BSTMap{}
}

func (m *BSTMap) GetSize() int {
	return m.size
}

func (m *BSTMap) IsEmpty() bool {
	return m.size == 0
}

func (m *BSTMap) Add(key interface{}, value interface{}) {
	m.root = m.add(m.root, key, value)
}

func (m *BSTMap) add(root *node, key interface{}, value interface{}) *node {
	if root == nil {
		m.size++
		return &node{key: key, value: value}
	}
	if utils.Compare(key, root.key) < 0 {
		root.left = m.add(root.left, key, value)
	} else if utils.Compare(key, root.key) > 0 {
		root.right = m.add(root.right, key, value)
	} else {
		root.value = value
	}
	return root
}

func (m *BSTMap) Set(key interface{}, value interface{}) {
	n := m.get(m.root, key)
	if n != nil {
		n.value = value
	}
}

func (m *BSTMap) Get(key interface{}) interface{} {
	n := m.get(m.root, key)
	if n == nil {
		return nil
	} else {
		return n.value
	}
}

func (m *BSTMap) get(root *node, key interface{}) *node {
	if root == nil {
		return nil
	}
	if utils.Compare(key, root.key) > 0 {
		return m.get(root.right, key)
	} else if utils.Compare(key, root.key) < 0 {
		return m.get(root.left, key)
	} else {
		return root
	}
}

func (m *BSTMap) Remove(key interface{}) {
	m.root = m.remove(m.root, key)
}

func (m *BSTMap) remove(root *node, key interface{}) *node {
	if root == nil {
		return nil
	}
	if utils.Compare(key, root.key) < 0 {
		root.left = m.remove(root.left, key)
		return root
	} else if utils.Compare(key, root.key) > 0 {
		root.right = m.remove(root.right, key)
		return root
	} else {
		leftChild, rightChild := root.left, root.right
		root.left, root.right = nil, nil
		if rightChild == nil {
			m.size--
			return leftChild
		}
		successor := m.minimum(rightChild)
		successor.right = m.removeMin(rightChild)
		successor.left = leftChild
		return successor
	}
}

func (m *BSTMap) minimum(root *node) *node {
	if root == nil || root.left == nil {
		return root
	}
	return m.minimum(root.left)
}

func (m *BSTMap) removeMin(root *node) *node {
	if root == nil {
		return nil
	}
	if root.left == nil {
		rightChild := root.right
		root.right = nil
		m.size--
		return rightChild
	} else {
		root.left = m.removeMin(root.left)
		return root
	}
}

func (m *BSTMap) Contains(key interface{}) bool {
	return m.contains(m.root, key)
}

func (m *BSTMap) contains(root *node, key interface{}) bool {
	if root == nil {
		return false
	}
	if utils.Compare(key, root.key) < 0 {
		return m.contains(root.left, key)
	} else if utils.Compare(key, root.key) > 0 {
		return m.contains(root.right, key)
	} else {
		return true
	}
}

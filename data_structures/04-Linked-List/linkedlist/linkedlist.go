package linkedlist

import (
	"LearnGo/data_structures/utils"
	"bytes"
	"fmt"
)

type node struct {
	value interface{}
	next  *node
}

func (n *node) String() string {
	return fmt.Sprintf("%v", n.value)
}

type LinkedList struct {
	dummyHead *node
	size      int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		dummyHead: &node{},
		size:      0,
	}
}

func (list *LinkedList) GetSize() int {
	return list.size
}

func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

// 在链表中不是一个常见操作，练习用:)
func (list *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > list.size {
		panic("Add failed, illegal index.")
	}
	prev := list.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = &node{
		value: e,
		next:  prev.next,
	}
	list.size++
}

func (list *LinkedList) AddFirst(e interface{}) {
	list.Add(0, e)
}

func (list *LinkedList) AddLast(e interface{}) {
	list.Add(list.size, e)
}

// 在链表中不是一个常见操作，练习用:)
func (list *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= list.size {
		panic("Get failed, illegal index.")
	}
	cur := list.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.value
}

func (list *LinkedList) GetFirst() interface{} {
	return list.Get(0)
}

func (list *LinkedList) GetLast() interface{} {
	return list.Get(list.size - 1)
}

func (list *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= list.size {
		panic("Set failed, illegal index.")
	}
	cur := list.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.value = e
}

func (list *LinkedList) Contains(e interface{}) bool {
	cur := list.dummyHead.next
	for cur != nil {
		if utils.Compare(cur.value, e) == 0 {
			return true
		}
		cur = cur.next
	}
	return false
}

// 在链表中不是一个常见操作，练习用:)
func (list *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= list.size {
		panic("Remove failed, illegal index.")
	}
	prev := list.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	delNode := prev.next
	prev.next = delNode.next
	delNode.next = nil
	list.size--
	return delNode.value
}

func (list *LinkedList) RemoveFirst() interface{} {
	return list.Remove(0)
}

func (list *LinkedList) RemoveLast() interface{} {
	return list.Remove(list.size - 1)
}

func (list *LinkedList) RemoveElement(e interface{}) bool {
	prev := list.dummyHead
	for prev.next != nil {
		if utils.Compare(prev.next.value, e) == 0 {
			break
		}
		prev = prev.next
	}
	if prev.next == nil {
		return false
	}
	delNode := prev.next
	prev.next = delNode.next
	delNode.next = nil
	list.size--
	return true
}

func (list *LinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := list.dummyHead.next
	for cur != nil {
		buffer.WriteString(fmt.Sprintf("%v -> ", cur.value))
		cur = cur.next
	}
	buffer.WriteString("nil")
	return buffer.String()
}

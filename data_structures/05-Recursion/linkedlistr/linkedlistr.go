package linkedlistr

import (
	"LearnGo/data_structures/utils"
	"bytes"
	"fmt"
)

type node struct {
	value interface{}
	next  *node
}

func (list *LinkedListR) add(head *node, index int, e interface{}) *node {
	if index == 0 {
		list.size++
		return &node{
			value: e,
			next:  head,
		}
	}
	head.next = list.add(head.next, index-1, e)
	return head
}

func (list *LinkedListR) get(head *node, index int) *node {
	if index == 0 {
		return head
	}
	return list.get(head.next, index-1)
}

func (list *LinkedListR) set(head *node, index int, e interface{}) {
	if index == 0 {
		head.value = e
		return
	}
	list.set(head.next, index-1, e)
}

func (list *LinkedListR) remove(head *node, index int) (*node, interface{}) {
	if index == 0 {
		next := head.next
		head.next = nil
		list.size--
		return next, head.value
	}
	next, delNodeValue := list.remove(head.next, index-1)
	head.next = next
	return head, delNodeValue
}

func (list *LinkedListR) contains(head *node, e interface{}) bool {
	if head == nil {
		return false
	}
	if utils.Compare(head.value, e) == 0 {
		return true
	}
	return list.contains(head.next, e)
}

func (list *LinkedListR) removeElement(head *node, e interface{}) (*node, bool) {
	if head == nil {
		return nil, false
	}
	if utils.Compare(head.value, e) == 0 {
		next := head.next
		head.next = nil
		list.size--
		return next, true
	}
	next, deleted := list.removeElement(head.next, e)
	head.next = next
	return head, deleted
}

type LinkedListR struct {
	head *node
	size int
}

func NewLinkedListR() *LinkedListR {
	return &LinkedListR{}
}

func (list *LinkedListR) GetSize() int {
	return list.size
}

func (list *LinkedListR) IsEmpty() bool {
	return list.size == 0
}

func (list *LinkedListR) Add(index int, e interface{}) {
	if index < 0 || index > list.size {
		panic("Add failed, illegal index.")
	}
	list.head = list.add(list.head, index, e)
}

func (list *LinkedListR) AddFirst(e interface{}) {
	list.Add(0, e)
}

func (list *LinkedListR) AddLast(e interface{}) {
	list.Add(list.size, e)
}

func (list *LinkedListR) Get(index int) interface{} {
	if index < 0 || index >= list.size {
		panic("Get failed, illegal index.")
	}
	return list.get(list.head, index).value
}

func (list *LinkedListR) GetFirst() interface{} {
	return list.Get(0)
}

func (list *LinkedListR) GetLast() interface{} {
	return list.Get(list.size - 1)
}

func (list *LinkedListR) Set(index int, e interface{}) {
	if index < 0 || index >= list.size {
		panic("Set failed, illegal index.")
	}
	list.set(list.head, index, e)
}

func (list *LinkedListR) Remove(index int) interface{} {
	if index < 0 || index >= list.size {
		panic("Remove failed, illegal index.")
	}
	newHead, delNodeValue := list.remove(list.head, index)
	list.head = newHead
	return delNodeValue
}

func (list *LinkedListR) RemoveFirst() interface{} {
	return list.Remove(0)
}

func (list *LinkedListR) RemoveLast() interface{} {
	return list.Remove(list.size - 1)
}

func (list *LinkedListR) Contains(e interface{}) bool {
	return list.contains(list.head, e)
}

func (list *LinkedListR) RemoveElement(e interface{}) bool {
	newHead, deleted := list.removeElement(list.head, e)
	list.head = newHead
	return deleted
}

func (list *LinkedListR) String() string {
	var buffer bytes.Buffer
	cur := list.head
	for cur != nil {
		buffer.WriteString(fmt.Sprintf("%v -> ", cur.value))
		cur = cur.next
	}
	buffer.WriteString("nil")
	return buffer.String()
}

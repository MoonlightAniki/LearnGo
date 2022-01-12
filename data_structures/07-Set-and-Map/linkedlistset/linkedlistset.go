package linkedlistset

import "LearnGo/data_structures/04-Linked-List/linkedlist"

type LinkedListSet struct {
	list *linkedlist.LinkedList
}

func NewLinkedListSet() *LinkedListSet {
	return &LinkedListSet{list: linkedlist.NewLinkedList()}
}

func (set *LinkedListSet) GetSize() int {
	return set.list.GetSize()
}

func (set *LinkedListSet) IsEmpty() bool {
	return set.list.IsEmpty()
}

func (set *LinkedListSet) Add(e interface{}) {
	if !set.Contains(e) {
		set.list.AddFirst(e)
	}
}

func (set *LinkedListSet) Remove(e interface{}) bool {
	return set.list.RemoveElement(e)
}

func (set *LinkedListSet) Contains(e interface{}) bool {
	return set.list.Contains(e)
}

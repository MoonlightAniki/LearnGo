package linkedlistqueue

import (
	"LearnGo/data_structures/04-Linked-List/linkedlist"
	"bytes"
	"fmt"
)

type LinkedListQueue struct {
	linkedList *linkedlist.LinkedList
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		linkedList: linkedlist.NewLinkedList(),
	}
}

func (queue *LinkedListQueue) Offer(e interface{}) {
	queue.linkedList.AddLast(e)
}

func (queue *LinkedListQueue) Poll() interface{} {
	return queue.linkedList.RemoveFirst()
}

func (queue *LinkedListQueue) Peek() interface{} {
	return queue.linkedList.GetFirst()
}

func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.linkedList.IsEmpty()
}

func (queue *LinkedListQueue) GetSize() int {
	return queue.linkedList.GetSize()
}

func (queue *LinkedListQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("LinkedListQueue: size=%d\n", queue.GetSize()))
	buffer.WriteString("front [")
	for i := 0; i < queue.GetSize(); i++ {
		buffer.WriteString(fmt.Sprintf("%v", queue.linkedList.Get(i)))
		if i != queue.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")
	return buffer.String()
}

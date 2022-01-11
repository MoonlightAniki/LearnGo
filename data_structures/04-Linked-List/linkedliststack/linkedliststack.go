package linkedliststack

import (
	"LearnGo/data_structures/04-Linked-List/linkedlist"
	"bytes"
	"fmt"
)

type LinkedListStack struct {
	linkedList *linkedlist.LinkedList
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		linkedList: linkedlist.NewLinkedList(),
	}
}

func (stack *LinkedListStack) Push(e interface{}) {
	stack.linkedList.AddFirst(e)
}

func (stack *LinkedListStack) Pop() interface{} {
	return stack.linkedList.RemoveFirst()
}

func (stack *LinkedListStack) Peek() interface{} {
	return stack.linkedList.GetFirst()
}

func (stack *LinkedListStack) IsEmpty() bool {
	return stack.linkedList.IsEmpty()
}

func (stack *LinkedListStack) GetSize() int {
	return stack.linkedList.GetSize()
}

func (stack *LinkedListStack) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("LinkedListStack: size=%d\n", stack.GetSize()))
	buffer.WriteString("top [")
	for i := 0; i < stack.GetSize(); i++ {
		buffer.WriteString(fmt.Sprintf("%v", stack.linkedList.Get(i)))
		if i != stack.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

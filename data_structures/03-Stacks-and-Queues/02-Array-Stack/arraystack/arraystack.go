package arraystack

import (
	"bytes"
	"fmt"
)

type ArrayStack []interface{}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{}
}

func (stack *ArrayStack) Push(e interface{}) {
	*stack = append(*stack, e)
}

func (stack *ArrayStack) Pop() interface{} {
	size := len(*stack)
	if size == 0 {
		panic("Pop failed, stack is empty.")
	}
	top := (*stack)[size-1]
	*stack = (*stack)[:size-1]
	return top
}

func (stack *ArrayStack) Peek() interface{} {
	size := len(*stack)
	if size == 0 {
		panic("Peek failed, stack is empty.")
	}
	return (*stack)[size-1]
}

func (stack *ArrayStack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *ArrayStack) GetSize() int {
	return len(*stack)
}

func (stack *ArrayStack) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("ArrayStack [")
	for i, e := range *stack {
		buffer.WriteString(fmt.Sprintf("%v", e))
		if i != len(*stack)-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")
	return buffer.String()
}

package arrayqueue

import (
	"bytes"
	"fmt"
)

type ArrayQueue []interface{}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{}
}

func (a *ArrayQueue) Offer(e interface{}) {
	*a = append(*a, e)
}

func (a *ArrayQueue) Poll() interface{} {
	if len(*a) == 0 {
		panic("Poll failed, queue is empty.")
	}
	front := (*a)[0]
	*a = (*a)[1:]
	return front
}

func (a *ArrayQueue) Peek() interface{} {
	if len(*a) == 0 {
		panic("Peek failed, queue is empty.")
	}
	return (*a)[0]
}

func (a *ArrayQueue) IsEmpty() bool {
	return len(*a) == 0
}

func (a *ArrayQueue) GetSize() int {
	return len(*a)
}

func (a *ArrayQueue) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("ArrayQueue: size=%d, capacity=%d\n", len(*a), cap(*a)))
	buffer.WriteString("front [")
	for i, e := range *a {
		buffer.WriteString(fmt.Sprintf("%v", e))
		if i != len(*a)-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")
	return buffer.String()
}

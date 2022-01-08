package loopqueue

import (
	"bytes"
	"fmt"
)

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func NewLoopQueue() *LoopQueue {
	// 默认容量设置为10
	capacity := 10
	return &LoopQueue{
		// data 空出一个位置用于区分队列是空的还是满的
		data:  make([]interface{}, capacity+1),
		front: 0,
		tail:  0,
		size:  0,
	}
}

func (q *LoopQueue) Offer(e interface{}) {
	// 如果队列是满的, 进行扩容
	if (q.tail+1)%len(q.data) == q.front {
		q.resize(q.GetCapacity() * 2)
	}
	q.data[q.tail] = e
	q.tail = (q.tail + 1) % len(q.data)
	q.size++
}

func (q *LoopQueue) Poll() interface{} {
	if q.IsEmpty() {
		panic("Poll failed, queue is empty.")
	}
	front := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % len(q.data)
	q.size--

	if q.size <= q.GetCapacity()/4 && q.GetCapacity()/2 > 0 {
		q.resize(q.GetCapacity() / 2)
	}
	return front
}

func (q *LoopQueue) Peek() interface{} {
	if q.IsEmpty() {
		panic("Peek failed, queue is empty.")
	}
	return q.data[q.front]
}

func (q *LoopQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LoopQueue) GetSize() int {
	return q.size
}

func (q *LoopQueue) GetCapacity() int {
	return len(q.data) - 1
}

func (q *LoopQueue) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("LoopQueue: size=%d, capacity=%d\n", q.size, q.GetCapacity()))
	buffer.WriteString("front [")
	for i := q.front; i != q.tail; i = (i + 1) % len(q.data) {
		buffer.WriteString(fmt.Sprintf("%v", q.data[i]))
		if (i+1)%len(q.data) != q.tail {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")
	return buffer.String()
}

func (q *LoopQueue) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity+1)
	for i := 0; i < q.size; i++ {
		newData[i] = q.data[(q.front+i)%len(q.data)]
	}
	q.data = newData
	q.front = 0
	q.tail = q.size
}

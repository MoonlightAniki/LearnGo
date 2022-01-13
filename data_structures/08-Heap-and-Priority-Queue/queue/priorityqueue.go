package queue

import "LearnGo/data_structures/08-Heap-and-Priority-Queue/heap"

type PriorityQueue struct {
	heap *heap.MaxHeap
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{heap.NewMaxHeap()}
}

func NewPriorityQueueWithArray(arr []interface{}) *PriorityQueue {
	return &PriorityQueue{heap.NewMaxHeapWithArray(arr)}
}

func (pq *PriorityQueue) Offer(e interface{}) {
	pq.heap.Add(e)
}

func (pq *PriorityQueue) Poll() interface{} {
	return pq.heap.ExtractMax()
}

func (pq *PriorityQueue) Peek() interface{} {
	return pq.heap.FindMax()
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.heap.IsEmpty()
}

func (pq *PriorityQueue) GetSize() int {
	return pq.heap.GetSize()
}

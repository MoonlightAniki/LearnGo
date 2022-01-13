package heap

import "LearnGo/data_structures/utils"

type MaxHeap struct {
	data []interface{}
	size int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func NewMaxHeapWithArray(arr []interface{}) *MaxHeap {
	h := &MaxHeap{
		data: arr,
		size: len(arr),
	}
	h.heapify()
	return h
}

func (h *MaxHeap) GetSize() int {
	return h.size
}

func (h *MaxHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *MaxHeap) Add(e interface{}) {
	h.data = append(h.data, e)
	h.size++
	h.shiftUp(h.size - 1)
}

func (h *MaxHeap) FindMax() interface{} {
	if h.IsEmpty() {
		panic("FindMax failed, MaxHeap is empty.")
	}
	return h.data[0]
}

func (h *MaxHeap) ExtractMax() interface{} {
	if h.IsEmpty() {
		panic("ExtractMax failed, MaxHeap is emtpy.")
	}
	max := h.data[0]
	h.data[0] = h.data[h.size-1]
	h.data[h.size-1] = nil
	h.size--
	h.shiftDown(0)
	return max
}

func (h *MaxHeap) shiftUp(k int) {
	// 存在父节点 && 父节点的值小于当前节点
	for parent(k) >= 0 && utils.Compare(h.data[parent(k)], h.data[k]) < 0 {
		h.swap(parent(k), k)
		k = parent(k)
	}
}

func (h *MaxHeap) shiftDown(k int) {
	for leftChild(k) < h.size {
		j := leftChild(k)
		if rightChild(k) < h.size && utils.Compare(h.data[rightChild(k)], h.data[leftChild(k)]) > 0 {
			j = rightChild(k)
		}
		if utils.Compare(h.data[k], h.data[j]) >= 0 {
			break
		}
		h.swap(k, j)
		k = j
	}
}

func (h *MaxHeap) heapify() {
	// 从最后一个非叶子节点开始执行 shiftDown
	for i := parent(h.size - 1); i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

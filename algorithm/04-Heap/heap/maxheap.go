package heap

import "LearnGo/data_structures/utils"

type MaxHeap struct {
	data     []interface{}
	capacity int
	count    int
}

func NewMaxHeap(capacity int) *MaxHeap {
	return &MaxHeap{
		data:     make([]interface{}, capacity+1),
		count:    0,
		capacity: capacity,
	}
}

func NewMaxHeapWithIntArray(arr []int) *MaxHeap {
	data := make([]interface{}, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		data[i+1] = arr[i]
	}
	maxHeap := &MaxHeap{
		data:     data,
		count:    len(arr),
		capacity: len(arr),
	}
	maxHeap.heapify()
	return maxHeap
}

func (h *MaxHeap) IsEmpty() bool {
	return h.count == 0
}

func (h *MaxHeap) GetSize() int {
	return h.count
}

func (h *MaxHeap) GetCapacity() int {
	return h.capacity
}

func (h *MaxHeap) Insert(e interface{}) {
	if h.count+1 > h.capacity {
		panic("Insert failed, heap is full.")
	}
	h.data[h.count+1] = e
	h.count++
	h.shiftUp(h.count)
}

func (h *MaxHeap) ExtractMax() interface{} {
	if h.IsEmpty() {
		panic("ExtractMax failed, heap is empty.")
	}
	max := h.data[1]
	h.data[1] = h.data[h.count]
	h.data[h.count] = nil
	h.count--
	h.shiftDown(1)
	return max
}

func (h *MaxHeap) shiftUp(k int) {
	// 存在父节点 && 父节点的值小于当前节点的值
	for k/2 >= 1 && utils.Compare(h.data[k/2], h.data[k]) < 0 {
		h.data[k/2], h.data[k] = h.data[k], h.data[k/2]
		k /= 2
	}
}

func (h *MaxHeap) shiftDown(k int) {
	// 存在左节点
	for k*2 <= h.count {
		j := k * 2 // data[j] 为此轮循环中与data[k]交换的元素
		// 存在右节点 && 右节点大于左节点
		if j+1 <= h.count && utils.Compare(h.data[j+1], h.data[j]) > 0 {
			j++
		}
		if utils.Compare(h.data[k], h.data[j]) >= 0 {
			break
		}
		h.data[k], h.data[j] = h.data[j], h.data[k]
		k = j
	}
}

func (h *MaxHeap) heapify() {
	// 从最后一个非叶子节点开始执行shiftDown操作
	for i := h.count / 2; i >= 1; i-- {
		h.shiftDown(i)
	}
}

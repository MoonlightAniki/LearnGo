package heap

type MinHeap struct {
	data []int
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.data[i] < h.data[j]
}

func (h *MinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	n := len(h.data)
	item := h.data[n-1]
	h.data = h.data[:n-1]
	return item
}

type MaxHeap struct {
	MinHeap
}

func (h *MaxHeap) Less(i int, j int) bool {
	return h.data[i] > h.data[j]
}

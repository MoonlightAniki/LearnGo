package heapsort

import "LearnGo/algorithm/04-Heap/heap"

func HeapSort1(arr []int) {
	maxHeap := heap.NewMaxHeap(len(arr))
	for i := 0; i < len(arr); i++ {
		maxHeap.Insert(arr[i])
	}
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = maxHeap.ExtractMax().(int)
	}
}

func HeapSort2(arr []int) {
	maxHeap := heap.NewMaxHeapWithIntArray(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = maxHeap.ExtractMax().(int)
	}
}

// 原地堆排序
func HeapSort3(arr []int) {
	// 从最后一个非叶子节点开始执行shiftDown操作
	// 最后一个节点的索引 len(arr)-1 => 最后一个非叶子节点的索引 (len(arr)-1 - 1) / 2
	for i := len(arr)/2 - 1; i >= 0; i-- {
		shiftDown(arr, len(arr)-1, i)
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		shiftDown(arr, i-1, 0)
	}
}

// 在 arr[0...r] 范围内执行 shiftDown
func shiftDown(arr []int, r int, k int) {
	for k*2+1 <= r {
		j := k*2 + 1
		if j+1 <= r && arr[j+1] > arr[j] {
			j++
		}
		if arr[k] >= arr[j] {
			break
		}
		arr[k], arr[j] = arr[j], arr[k]
		k = j
	}
}

// 相对 HeapSort3 速度并没有多少提升
func HeapSort4(arr []int) {
	// 从最后一个非叶子节点开始执行shiftDown操作
	// 最后一个节点的索引 len(arr)-1 => 最后一个非叶子节点的索引 (len(arr)-1 - 1) / 2
	for i := len(arr)/2 - 1; i >= 0; i-- {
		shiftDown2(arr, len(arr)-1, i)
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		shiftDown2(arr, i-1, 0)
	}
}

func shiftDown2(arr []int, r int, k int) {
	v := arr[k]
	for 2*k+1 <= r {
		j := 2*k + 1
		if j+1 <= r && arr[j+1] > arr[j] {
			j++
		}
		if v >= arr[j] {
			break
		}
		arr[k] = arr[j]
		k = j
	}
	arr[k] = v
}

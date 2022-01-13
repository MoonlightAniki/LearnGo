package main

import (
	"LearnGo/algorithm/02-Sorting-Algorithms/sorttesthelper"
	"LearnGo/algorithm/04-Heap/heapsort"
	myheap "LearnGo/data_structures/08-Heap-and-Priority-Queue/heaptest/heap"
	"container/heap"
)

func main() {
	n := 1000000
	arr := sorttesthelper.GenerateRandomArray(n, 0, n)
	arr1 := sorttesthelper.CopyIntArray(arr)
	arr2 := sorttesthelper.CopyIntArray(arr)
	arr3 := sorttesthelper.CopyIntArray(arr)
	arr4 := sorttesthelper.CopyIntArray(arr)

	sorttesthelper.TestSort("HeapSort", heapSort, arr)
	sorttesthelper.TestSort("HeapSort1", heapsort.HeapSort1, arr1)
	sorttesthelper.TestSort("HeapSort2", heapsort.HeapSort2, arr2)
	sorttesthelper.TestSort("HeapSort3", heapsort.HeapSort3, arr3)
	sorttesthelper.TestSort("HeapSort4", heapsort.HeapSort4, arr4)
}

func heapSort(arr []int) {
	maxHeap := &myheap.MaxHeap{}
	for i := 0; i < len(arr); i++ {
		heap.Push(maxHeap, arr[i])
	}
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = heap.Pop(maxHeap).(int)
	}
}

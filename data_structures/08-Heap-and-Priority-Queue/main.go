package main

import (
	"LearnGo/algorithm/02-Sorting-Algorithms/sorttesthelper"
	"LearnGo/data_structures/08-Heap-and-Priority-Queue/heap"
	"fmt"
	"math/rand"
)

func heapSort1(arr []int) {
	h := heap.NewMaxHeap()
	for i := 0; i < len(arr); i++ {
		h.Add(arr[i])
	}
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = h.ExtractMax().(int)
	}
}

func heapSort2(arr []int) {
	data := make([]interface{}, len(arr))
	for i := 0; i < len(arr); i++ {
		data[i] = arr[i]
	}
	h := heap.NewMaxHeapWithArray(data)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = h.ExtractMax().(int)
	}
}

func heapSort3(arr []int) {
	for i := (len(arr) - 1 - 1) / 2; i >= 0; i-- {
		shiftDown(arr, i, len(arr)-1)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		shiftDown(arr, 0, i-1)
	}
}

func shiftDown(arr []int, k int, r int) {
	for 2*k+1 <= r {
		j := 2*k + 1
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

func main() {
	//test1()
	//testHeapSort()
}

func testHeapSort() {
	n := 1000000
	arr := sorttesthelper.GenerateRandomArray(n, 0, n)
	arr1 := sorttesthelper.CopyIntArray(arr)
	arr2 := sorttesthelper.CopyIntArray(arr)

	sorttesthelper.TestSort("HeapSort1", heapSort1, arr)
	sorttesthelper.TestSort("HeapSort2", heapSort2, arr1)
	sorttesthelper.TestSort("HeapSort3", heapSort3, arr2)
}

func test1() {
	h := heap.NewMaxHeap()
	for i := 0; i < 100; i++ {
		h.Add(rand.Intn(100))
	}
	for !h.IsEmpty() {
		fmt.Printf("%v ", h.ExtractMax())
	}
}

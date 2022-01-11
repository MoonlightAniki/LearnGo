package main

import (
	"LearnGo/algorithm/02-Sorting-Algorithms/mergeSort"
	"LearnGo/algorithm/02-Sorting-Algorithms/quickSort"
	"LearnGo/algorithm/02-Sorting-Algorithms/shellSort"
	"LearnGo/algorithm/02-Sorting-Algorithms/sortTestHelper"
	"LearnGo/algorithm/04-Heap/heapsort"
)

func main() {
	n := 500000
	arr := sortTestHelper.GenerateRandomArray(n, 0, n) // 随机数组
	//arr := sortTestHelper.GenerateNearlyOrderedArray(n, 100) // 近乎有序的数组
	//arr := sortTestHelper.GenerateRandomArray(n, 0, 10) // 有大量重复元素的数组
	//arr2 := sortTestHelper.CopyIntArray(arr)
	//arr3 := sortTestHelper.CopyIntArray(arr)
	//arr4 := sortTestHelper.CopyIntArray(arr)
	//arr5 := sortTestHelper.CopyIntArray(arr)
	//arr6 := sortTestHelper.CopyIntArray(arr)
	arr7 := sortTestHelper.CopyIntArray(arr)
	arr8 := sortTestHelper.CopyIntArray(arr)
	//arr9 := sortTestHelper.CopyIntArray(arr)
	//arr10 := sortTestHelper.CopyIntArray(arr)
	arr11 := sortTestHelper.CopyIntArray(arr)
	arr12 := sortTestHelper.CopyIntArray(arr)
	arr13 := sortTestHelper.CopyIntArray(arr)
	arr14 := sortTestHelper.CopyIntArray(arr)
	arr15 := sortTestHelper.CopyIntArray(arr)
	arr16 := sortTestHelper.CopyIntArray(arr)
	arr17 := sortTestHelper.CopyIntArray(arr)
	arr18 := sortTestHelper.CopyIntArray(arr)
	arr19 := sortTestHelper.CopyIntArray(arr)
	arr20 := sortTestHelper.CopyIntArray(arr)
	arr21 := sortTestHelper.CopyIntArray(arr)
	arr22 := sortTestHelper.CopyIntArray(arr)

	//sortTestHelper.TestSort("SelectionSort", selectionSort.SelectionSort, arr)
	//sortTestHelper.TestSort("InsertionSort", insertionSort.InsertionSort, arr2)
	//sortTestHelper.TestSort("InsertionSort2", insertionSort.InsertionSort2, arr3)
	//sortTestHelper.TestSort("BubbleSort", bubbleSort.BubbleSort, arr4)
	//sortTestHelper.TestSort("BubbleSort2", bubbleSort.BubbleSort2, arr5)
	//sortTestHelper.TestSort("BubbleSort3", bubbleSort.BubbleSort3, arr6)
	sortTestHelper.TestSort("shellSort", shellSort.ShellSort, arr7)
	sortTestHelper.TestSort("shellSort2", shellSort.ShellSort2, arr8)
	// mergeSort 和 mergeSortBU 由于在排序过程中创建了大量的辅助数组，所以速度比较慢，甚至比选择排序还要慢...
	//sortTestHelper.TestSort("mergeSort", mergeSort.MergeSort, arr9)
	//sortTestHelper.TestSort("mergeSortBU", mergeSort.MergeSortBU, arr10)
	sortTestHelper.TestSort("mergeSort2", mergeSort.MergeSort2, arr11)
	sortTestHelper.TestSort("mergeSortBU2", mergeSort.MergeSortBU2, arr12)
	sortTestHelper.TestSort("mergeSort3", mergeSort.MergeSort3, arr13)
	sortTestHelper.TestSort("mergeSortBU3", mergeSort.MergeSortBU3, arr14)
	sortTestHelper.TestSort("quickSort", quickSort.QuickSort, arr15)
	sortTestHelper.TestSort("quickSort2", quickSort.QuickSort2, arr16)
	sortTestHelper.TestSort("quickSort3", quickSort.QuickSort3, arr17)
	sortTestHelper.TestSort("quickSort2Ways", quickSort.QuickSort2Ways, arr18)
	sortTestHelper.TestSort("quickSort3Ways", quickSort.QuickSort3Ways, arr19)
	sortTestHelper.TestSort("HeapSort1", heapsort.HeapSort1, arr20)
	sortTestHelper.TestSort("HeapSort2", heapsort.HeapSort2, arr21)
	sortTestHelper.TestSort("HeapSort3", heapsort.HeapSort3, arr22)

}

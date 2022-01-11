package main

import (
	"LearnGo/algorithm/02-Sorting-Algorithms/mergesort"
	"LearnGo/algorithm/02-Sorting-Algorithms/quicksort"
	"LearnGo/algorithm/02-Sorting-Algorithms/shellsort"
	"LearnGo/algorithm/02-Sorting-Algorithms/sorttesthelper"
	"LearnGo/algorithm/04-Heap/heapsort"
)

func main() {
	n := 500000
	arr := sorttesthelper.GenerateRandomArray(n, 0, n) // 随机数组
	//arr := sortTestHelper.GenerateNearlyOrderedArray(n, 100) // 近乎有序的数组
	//arr := sortTestHelper.GenerateRandomArray(n, 0, 10) // 有大量重复元素的数组
	//arr2 := sorttesthelper.CopyIntArray(arr)
	//arr3 := sorttesthelper.CopyIntArray(arr)
	//arr4 := sorttesthelper.CopyIntArray(arr)
	//arr5 := sorttesthelper.CopyIntArray(arr)
	//arr6 := sorttesthelper.CopyIntArray(arr)
	arr7 := sorttesthelper.CopyIntArray(arr)
	arr8 := sorttesthelper.CopyIntArray(arr)
	//arr9 := sorttesthelper.CopyIntArray(arr)
	//arr10 := sorttesthelper.CopyIntArray(arr)
	arr11 := sorttesthelper.CopyIntArray(arr)
	arr12 := sorttesthelper.CopyIntArray(arr)
	arr13 := sorttesthelper.CopyIntArray(arr)
	arr14 := sorttesthelper.CopyIntArray(arr)
	arr15 := sorttesthelper.CopyIntArray(arr)
	arr16 := sorttesthelper.CopyIntArray(arr)
	arr17 := sorttesthelper.CopyIntArray(arr)
	arr18 := sorttesthelper.CopyIntArray(arr)
	arr19 := sorttesthelper.CopyIntArray(arr)
	arr20 := sorttesthelper.CopyIntArray(arr)
	arr21 := sorttesthelper.CopyIntArray(arr)
	arr22 := sorttesthelper.CopyIntArray(arr)

	//sorttesthelper.TestSort("SelectionSort", selectionsort.SelectionSort, arr)
	//sorttesthelper.TestSort("InsertionSort", insertionsort.InsertionSort, arr2)
	//sorttesthelper.TestSort("InsertionSort2", insertionsort.InsertionSort2, arr3)
	//sorttesthelper.TestSort("BubbleSort", bubblesort.BubbleSort, arr4)
	//sorttesthelper.TestSort("BubbleSort2", bubblesort.BubbleSort2, arr5)
	//sorttesthelper.TestSort("BubbleSort3", bubblesort.BubbleSort3, arr6)
	sorttesthelper.TestSort("shellSort", shellsort.ShellSort, arr7)
	sorttesthelper.TestSort("shellSort2", shellsort.ShellSort2, arr8)
	// mergeSort 和 mergeSortBU 由于在排序过程中创建了大量的辅助数组，所以速度比较慢，甚至比选择排序还要慢...
	//sorttesthelper.TestSort("mergeSort", mergesort.MergeSort, arr9)
	//sorttesthelper.TestSort("mergeSortBU", mergesort.MergeSortBU, arr10)
	sorttesthelper.TestSort("mergeSort2", mergesort.MergeSort2, arr11)
	sorttesthelper.TestSort("mergeSortBU2", mergesort.MergeSortBU2, arr12)
	sorttesthelper.TestSort("mergeSort3", mergesort.MergeSort3, arr13)
	sorttesthelper.TestSort("mergeSortBU3", mergesort.MergeSortBU3, arr14)
	sorttesthelper.TestSort("quickSort", quicksort.QuickSort, arr15)
	sorttesthelper.TestSort("quickSort2", quicksort.QuickSort2, arr16)
	sorttesthelper.TestSort("quickSort3", quicksort.QuickSort3, arr17)
	sorttesthelper.TestSort("quickSort2Ways", quicksort.QuickSort2Ways, arr18)
	sorttesthelper.TestSort("quickSort3Ways", quicksort.QuickSort3Ways, arr19)
	sorttesthelper.TestSort("HeapSort1", heapsort.HeapSort1, arr20)
	sorttesthelper.TestSort("HeapSort2", heapsort.HeapSort2, arr21)
	sorttesthelper.TestSort("HeapSort3", heapsort.HeapSort3, arr22)

}

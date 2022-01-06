package mergeSort

import "LearnGo/algorithm/02-Sorting-Algorithms/insertionSort"

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

// 对 arr[l...r]范围内的元素进行归并排序
func mergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	if arr[mid] > arr[mid+1] {
		merge(arr, l, mid, r)
	}
}

func MergeSortBU(arr []int) {
	n := len(arr)
	for size := 1; size < n; size += size {
		for i := 0; i+size /*保证第二部分包含元素*/ < n; i += 2 * size {
			// 对 arr[i...i+size-1]、arr[i+size...i+2*size-1]
			if arr[i+size-1] > arr[i+size] {
				merge(arr, i, i+size-1, minOf(n-1, i+2*size-1))
			}
		}
	}
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 对 arr[l...mid]、arr[mid+1...r] 两部分进行归并操作
func merge(arr []int, l int, mid int, r int) {
	aux := make([]int, len(arr))
	for i := l; i <= r; i++ {
		aux[i-l] = arr[i]
	}
	i := l
	j := mid + 1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}
}

// 整个归并排序过程中只创建一个辅助数组aux
func MergeSort2(arr []int) {
	aux := make([]int, len(arr))
	mergeSort2(arr, 0, len(arr)-1, aux)
}

func mergeSort2(arr []int, l int, r int, aux []int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeSort2(arr, l, mid, aux)
	mergeSort2(arr, mid+1, r, aux)
	if arr[mid] > arr[mid+1] {
		merge2(arr, l, mid, r, aux)
	}
}

func merge2(arr []int, l int, mid int, r int, aux []int) {
	for i := l; i <= r; i++ {
		aux[i] = arr[i]
	}
	i := l
	j := mid + 1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = aux[j]
			j++
		} else if j > r {
			arr[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			arr[k] = aux[i]
			i++
		} else {
			arr[k] = aux[j]
			j++
		}
	}
}

func MergeSortBU2(arr []int) {
	n := len(arr)
	aux := make([]int, n)
	for size := 1; size < n; size += size {
		for i := 0; i+size < n; i += 2 * size {
			if arr[i+size-1] > arr[i+size] {
				merge2(arr, i, i+size-1, minOf(n-1, i+2*size-1), aux)
			}
		}
	}
}

func MergeSort3(arr []int) {
	n := len(arr)
	aux := make([]int, n)
	mergeSort3(arr, 0, n-1, aux)
}

func mergeSort3(arr []int, l int, r int, aux []int) {
	if r-l+1 <= 16 {
		insertionSort.InsertionSort3(arr, l, r)
		return
	}
	mid := l + (r-l)/2
	mergeSort3(arr, l, mid, aux)
	mergeSort3(arr, mid+1, r, aux)
	merge3(arr, l, mid, r, aux)
}

func merge3(arr []int, l int, mid int, r int, aux []int) {
	merge2(arr, l, mid, r, aux)
}

func MergeSortBU3(arr []int) {
	n := len(arr)
	size := 16
	for i := 0; i < n; i += size {
		// 对 arr[i...i+size-1] 范围内元素进行插入排序
		insertionSort.InsertionSort3(arr, i, minOf(n-1, i+size-1))
	}
	aux := make([]int, n)
	for ; size < n; size += size {
		for i := 0; i+size < n; i += 2 * size {
			if arr[i+size-1] > arr[i+size] {
				merge3(arr, i, i+size-1, minOf(n-1, i+2*size-1), aux)
			}
		}
	}
}

package quickSort

import (
	"LearnGo/algorithm/02-Sorting-Algorithms/insertionSort"
	"math/rand"
)

// region 快速排序 第一个版本
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

// 对 arr[l...r] 范围内元素进行快速排序
func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

// partition 完成后返回 p, 满足 arr[l...p-1] < arr[p] && arr[p+1...r] > arr[p]
func partition(arr []int, l int, r int) int {
	v := arr[l]
	j := l // arr[l+1...j] < v, arr[j+1...i) > v
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// endregion

// region 快速排序 第二个版本
// partition 操作时随机选择一个元素的值作为标定点 pivot
func QuickSort2(arr []int) {
	quickSort2(arr, 0, len(arr)-1)
}

func quickSort2(arr []int, l int, r int) {
	if l >= r {
		return
	}
	p := partition2(arr, l, r)
	quickSort2(arr, l, p-1)
	quickSort2(arr, p+1, r)
}

func partition2(arr []int, l int, r int) int {
	x := l + rand.Intn(r-l+1)
	arr[l], arr[x] = arr[x], arr[l]

	v := arr[l]
	j := l // arr[l+1...j] < v, arr[j+1...i) > v
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// endregion

// region 快速排序 第三个版本
// 在第二个版本的基础上，小范围内使用插入排序进行优化
func QuickSort3(arr []int) {
	quickSort3(arr, 0, len(arr)-1)
}

func quickSort3(arr []int, l int, r int) {
	if r-l+1 <= 16 {
		insertionSort.InsertionSort3(arr, l, r)
		return
	}
	p := partition3(arr, l, r)
	quickSort3(arr, l, p-1)
	quickSort3(arr, p+1, r)
}

var partition3 = partition2

// endregion

// region 双路快排
func QuickSort2Ways(arr []int) {
	quickSort2Ways(arr, 0, len(arr)-1)
}

func quickSort2Ways(arr []int, l int, r int) {
	if r-l+1 <= 16 {
		insertionSort.InsertionSort3(arr, l, r)
		return
	}
	p := partition2Ways(arr, l, r)
	quickSort2Ways(arr, l, p-1)
	quickSort2Ways(arr, p+1, r)
}

func partition2Ways(arr []int, l int, r int) int {
	x := rand.Intn(r-l+1) + l
	arr[l], arr[x] = arr[x], arr[l]

	v := arr[l]
	// arr[l+1...i) <= v, arr(j...r] >= v
	i := l + 1
	j := r
	for {
		for i <= r && arr[i] < v {
			i++
		}
		for j >= l+1 && arr[j] > v {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// endregion

// region 三路快排
func QuickSort3Ways(arr []int) {
	quickSort3Ways(arr, 0, len(arr)-1)
}

func quickSort3Ways(arr []int, l int, r int) {
	if r-l+1 <= 16 {
		insertionSort.InsertionSort3(arr, l, r)
		return
	}
	lt, gt := partition3Ways(arr, l, r)
	quickSort3Ways(arr, l, lt)
	quickSort3Ways(arr, gt, r)
}

// 返回 lt, gt 满足 arr[l...lt] < v, arr[gt...r] > v
func partition3Ways(arr []int, l int, r int) (int, int) {
	x := rand.Intn(r-l+1) + l
	arr[l], arr[x] = arr[x], arr[l]

	v := arr[l]
	// arr[l+1...lt] < v, arr[gt...r] > v
	lt := l
	gt := r + 1
	for i := l + 1; i < gt; {
		if arr[i] < v {
			lt++
			arr[lt], arr[i] = arr[i], arr[lt]
			i++
		} else if arr[i] > v {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		} else {
			// arr[i] == v
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	lt--
	return lt, gt
}

// endregion

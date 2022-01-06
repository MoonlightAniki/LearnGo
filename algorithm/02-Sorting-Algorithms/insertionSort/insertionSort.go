package insertionSort

// 插入排序
func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j-1 >= 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

func InsertionSort2(arr []int) {
	for i := 0; i < len(arr); i++ {
		e := arr[i]
		// j 为 e 最终应该在的位置
		j := i
		for ; j-1 >= 0 && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}

// 对 arr[l...r] 范围内元素进行插入排序
func InsertionSort3(arr []int, l int, r int) {
	for i := l; i <= r; i++ {
		e := arr[i]
		j := i
		for ; j-1 >= l && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}

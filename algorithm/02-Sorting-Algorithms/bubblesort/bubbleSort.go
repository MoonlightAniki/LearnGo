package bubblesort

// 冒泡排序
func BubbleSort(arr []int) {
	n := len(arr)
	for n > 0 {
		for i := 0; i+1 < n; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		n--
	}
}

// 使用 swapped 提前结束循环
func BubbleSort2(arr []int) {
	n := len(arr)
	for {
		swapped := false
		for i := 0; i+1 < n; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		n--
		if !swapped {
			break
		}
	}
}

// 使用 newn 进行优化
func BubbleSort3(arr []int) {
	n := len(arr)
	for {
		newn := 0
		for i := 0; i+1 < n; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				newn = i + 1
			}
		}
		n = newn
		if n == 0 {
			break
		}
	}
}

// 使用 swapped 和 newn 进行优化之后，速度反而比未优化版本慢！

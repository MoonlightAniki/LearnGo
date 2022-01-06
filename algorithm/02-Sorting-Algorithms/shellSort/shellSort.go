package shellSort

// 希尔排序
// 需要保证最后一次巡环时h的值为1，否则无法保证排序后数组是有序的！
func ShellSort(arr []int) {
	n := len(arr)
	h := n / 2
	for h > 0 {
		for i := 0; i < n; i++ {
			e := arr[i]
			j := i
			for ; j-h >= 0 && arr[j-h] > e; j -= h {
				arr[j] = arr[j-h]
			}
			arr[j] = e
		}
		h /= 3
	}
}

func ShellSort2(arr []int) {
	n := len(arr)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h > 0 {
		for i := 0; i < n; i++ {
			e := arr[i]
			j := i
			for ; j-h >= 0 && arr[j-h] > e; j -= h {
				arr[j] = arr[j-h]
			}
			arr[j] = e
		}
		h /= 3
	}
}

package solution0048

func rotate(matrix [][]int) {
	for _, row := range matrix {
		reverse(row)
	}

	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			matrix[i][j], matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i], matrix[i][j]
		}
	}
}

func reverse(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

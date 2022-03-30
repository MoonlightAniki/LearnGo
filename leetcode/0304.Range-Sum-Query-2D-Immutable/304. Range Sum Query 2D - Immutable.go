package solution0304

// 执行用时： 616 ms, 在所有 Go 提交中击败了 38.23% 的用户
/*type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sums := make([][]int, m)
	for i := 0; i < m; i++ {
		sums[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			sums[i][j] = sums[i][j-1] + matrix[i][j-1]
		}
	}
	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.sums[i][col2+1] - this.sums[i][col1]
	}
	return sum
}*/

// 执行用时：592 ms, 在所有 Go 提交中击败了 53.09% 的用户
type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sums := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		sums[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sums[i+1][j+1] = sums[i+1][j] + (sums[i][j+1] - sums[i][j]) + matrix[i][j]
		}
	}
	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sums[row2+1][col2+1] - this.sums[row2+1][col1] - this.sums[row1][col2+1] + this.sums[row1][col1]
}

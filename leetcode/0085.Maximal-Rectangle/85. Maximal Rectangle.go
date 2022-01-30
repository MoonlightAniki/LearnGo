package solution0085

/*
Given a rows x cols binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.



Example 1:
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 6
Explanation: The maximal rectangle is shown in the above picture.


Example 2:
Input: matrix = [["0"]]
Output: 0


Example 3:
Input: matrix = [["1"]]
Output: 1


Constraints:
rows == matrix.length
cols == matrix[i].length
1 <= row, cols <= 200
matrix[i][j] is '0' or '1'.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-rectangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 执行用时： 16 ms, 在所有 Go 提交中击败了 16.97% 的用户
func maximalRectangle(matrix [][]byte) int {
	R, C := len(matrix), len(matrix[0])
	dp := make([][]int, R)
	for i := range dp {
		dp[i] = make([]int, C)
	}
	for i := 0; i < R; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
		for j := 1; j < C; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i][j-1] + 1
			}
		}
	}
	res := 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			w := dp[i][j]
			for k := i; k >= 0; k-- {
				w = minOf(w, dp[k][j])
				h := i - k + 1
				res = maxOf(res, w*h)
			}
		}
	}
	return res
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

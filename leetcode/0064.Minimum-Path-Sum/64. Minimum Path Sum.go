package solution0064

import "fmt"

/*
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.


Example 1:
Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.

Example 2:
Input: grid = [[1,2,3],[4,5,6]]
Output: 12


Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = grid[i][j] + minOf(dp[j], dp[j-1])
		}
	}
	return dp[n-1]
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Test() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
	fmt.Println(minPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}

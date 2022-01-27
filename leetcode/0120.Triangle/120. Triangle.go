package solution0120

import "fmt"

/*
Given a triangle array, return the minimum path sum from top to bottom.

For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.



Example 1:
Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
   2
  3 4
 6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).


Example 2:
Input: triangle = [[-10]]
Output: -10


Constraints:
1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-10^4 <= triangle[i][j] <= 10^4


Follow up: Could you do this using only O(n) extra space, where n is the total number of rows in the triangle?
*/

/*func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := i; j >= 0; j-- {
			if j == 0 {
				dp[j] = dp[j] + triangle[i][j]
			} else if j == i {
				dp[j] = dp[j-1] + triangle[i][j]
			} else {
				dp[j] = minOf(dp[j-1], dp[j]) + triangle[i][j]
			}
		}
	}
	res := dp[0]
	for i := 1; i < len(dp); i++ {
		res = minOf(res, dp[i])
	}
	return res
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}*/

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == i {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += minOf(triangle[i-1][j-1], triangle[i-1][j])
			}
		}
	}
	res := triangle[n-1][0]
	for j := 1; j < len(triangle[n-1]); j++ {
		res = minOf(res, triangle[n-1][j])
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

func Test() {
	fmt.Println(
		minimumTotal([][]int{
			{2},
			{3, 4},
			{6, 5, 7},
			{4, 1, 8, 3},
		}))
}

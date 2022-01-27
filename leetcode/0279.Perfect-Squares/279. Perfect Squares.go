package solution0279

/*
Given an integer n, return the least number of perfect square numbers that sum to n.

A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself. For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.


Example 1:
Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.


Example 2:
Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.


Constraints:
1 <= n <= 10^4
*/

// Runtime: 89 ms, faster than 40.49% of Go online submissions for Perfect Squares.
/*var memo []int

func numSquares(n int) int {
	memo = make([]int, n+1)
	return helper(n)
}

func helper(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	res := n
	for i := 1; i*i <= n; i++ {
		res = minOf(res, 1+helper(n-i*i))
	}
	memo[n] = res
	return res
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}*/

// Runtime: 31 ms, faster than 77.07% of Go online submissions for Perfect Squares.
func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		min := i
		for j := 1; j*j <= i; j++ {
			min = minOf(min, 1+dp[i-j*j])
		}
		dp[i] = min
	}
	return dp[n]
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

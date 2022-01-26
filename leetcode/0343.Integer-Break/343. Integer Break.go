package solution0343

import "fmt"

/*
Given an integer n, break it into the sum of k positive integers, where k >= 2, and maximize the product of those integers.

Return the maximum product you can get.



Example 1:
Input: n = 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.


Example 2:
Input: n = 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.


Constraints:
2 <= n <= 58
*/

/*var memo []int

// 记忆化搜索，时间复杂度 O(n^2)
func integerBreak(n int) int {
	memo = make([]int, n+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	return breakInteger(n)
}

// 返回将 n 拆成至少 2 个正整数，得到的最大乘积
func breakInteger(n int) int {
	if n == 1 {
		return 1
	}
	if memo[n] != -1 {
		return memo[n]
	}
	res := -1
	for i := 1; i < n; i++ {
		// 将 n 分成 i 和 n - i
		res = max3(res, i*(n-i), i*breakInteger(n-i))
	}
	memo[n] = res
	return res
}*/

// 动态规划, 时间复杂度 O(n^2)
func integerBreak(n int) int {
	// dp[i] 表示 将 i 分成至少两部分得到的最大乘积
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		max := -1
		for j := 1; j < i; j++ {
			// 将 i 分成 j 和 i - j 两部分, 或者将 i-j 继续拆分
			max = max3(max, j*(i-j), j*dp[i-j])
		}
		dp[i] = max
	}
	return dp[n]
}

func max2(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func max3(a, b, c int) int {
	return max2(a, max2(b, c))
}

func Test() {
	fmt.Println(integerBreak(10))
	fmt.Println(integerBreak(2))
}

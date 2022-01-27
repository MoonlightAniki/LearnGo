package solution0322

/*
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.



Example 1:
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1


Example 2:
Input: coins = [2], amount = 3
Output: -1


Example 3:
Input: coins = [1], amount = 0
Output: 0


Constraints:
1 <= coins.length <= 12
1 <= coins[i] <= 2^31 - 1
0 <= amount <= 10^4
*/

// Runtime: 35 ms, faster than 35.58% of Go online submissions for Coin Change.
/*var memo []int
var max_count int

func coinChange(coins []int, amount int) int {
	memo = make([]int, amount+1)
	max_count = amount + 1
	for i := 0; i < len(memo); i++ {
		memo[i] = max_count
	}
	return helper(coins, amount)
}

func helper(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if memo[amount] != max_count {
		return memo[amount]
	}
	res := max_count
	for _, c := range coins {
		if c <= amount {
			count := helper(coins, amount-c)
			if count != -1 {
				res = minOf(res, 1+count)
			}
		}
	}
	if res == max_count {
		res = -1
	}
	memo[amount] = res
	return res
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}*/

// Runtime: 8 ms, faster than 94.09% of Go online submissions for Coin Change.
/*func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := i + 1
		for _, c := range coins {
			if c <= i {
				count := dp[i-c]
				if count != -1 {
					min = minOf(min, 1+count)
				}
			}
		}
		if min == i+1 {
			min = -1
		}
		dp[i] = min
	}
	return dp[amount]
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}*/

// Runtime: 15 ms, faster than 62.23% of Go online submissions for Coin Change.
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for _, c := range coins {
		for j := c; j <= amount; j++ {
			dp[j] = minOf(dp[j], 1+dp[j-c])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

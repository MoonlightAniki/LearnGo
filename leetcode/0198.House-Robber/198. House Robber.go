package solution0198

import "fmt"

/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.



Example 1:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.


Example 2:
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.


Constraints:
1 <= nums.length <= 100
0 <= nums[i] <= 400
*/

// 记忆化搜索，时间复杂度 O(n)
/*var memo []int

func rob(nums []int) int {
	memo = make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	return tryRob(nums, 0)
}

// 考虑从 nums[index...n-1] 范围内偷取财物能获得的最大值
func tryRob(nums []int, index int) int {
	if index >= len(nums) {
		return 0
	}
	if memo[index] != -1 {
		return memo[index]
	}
	memo[index] = maxOf(
		tryRob(nums, index+1),             //不偷 nums[index]
		nums[index]+tryRob(nums, index+2), // 偷取 nums[index]
	)
	return memo[index]
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}*/

func rob(nums []int) int {
	// dp[i] 表示从 nums[0...i] 范围内偷取财物能获取的最大收益
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		// 不偷取 nums[i]
		a := dp[i-1]
		// 偷取 nums[i]
		b := nums[i]
		if i-2 >= 0 {
			b += dp[i-2]
		}
		dp[i] = maxOf(a, b)
	}
	return dp[len(nums)-1]
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Test() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

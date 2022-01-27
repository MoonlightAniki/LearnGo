package solution0213

import "fmt"

/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed.
All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile,
adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses
were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can
rob tonight without alerting the police.


Example 1:
Input: nums = [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.


Example 2:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.


Example 3:
Input: nums = [1,2,3]
Output: 3


Constraints:
1 <= nums.length <= 100
0 <= nums[i] <= 1000
*/
/*func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return maxOf(nums[0], nums[1])
	}
	return maxOf(tryRob(nums[1:]), tryRob(nums[:n-1]))
}

func tryRob(nums []int) int {
	n := len(nums)
	// dp[i] 表示从 nums[0...i]范围内偷取财物能获取的最大收益
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		// 不偷取 nums[i]
		a := dp[i-1]
		// 偷取 nums[i]
		b := nums[i]
		if i-2 >= 0 {
			b += dp[i-2]
		}
		dp[i] = maxOf(a, b)
	}
	return dp[n-1]
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}*/

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return maxOf(nums[0], nums[1])
	}
	return maxOf(tryRob(nums, 1, n-1), tryRob(nums, 0, n-2))
}

func tryRob(nums []int, start int, end int) int {
	prevMax, curMax := nums[start], maxOf(nums[start], nums[start+1])
	for i := start + 2; i <= end; i++ {
		temp := curMax
		curMax = maxOf(nums[i]+prevMax, curMax)
		prevMax = temp
	}
	return curMax
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Test() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{1}))
}

package solution0416

import "fmt"

/*
Given a non-empty array nums containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.



Example 1:
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].


Example 2:
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.


Constraints:
1 <= nums.length <= 200
1 <= nums[i] <= 100
*/

// 记忆化搜索
/*var memo [][]int

func canPartition(nums []int) bool {
	// 问题可以转化为从 nums 中取出若干个元素，使得这些元素的和等于 sum/2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	n := len(nums)
	c := sum / 2
	memo = make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, c+1)
		for j := 0; j <= c; j++ {
			memo[i][j] = -1
		}
	}
	return tryPartition(nums, n-1, c)
}

// 从 nums[0...index] 从取出若干个物品放入容量为c的背包中
func tryPartition(nums []int, index int, c int) bool {
	if index < 0 {
		return false
	}
	if c == 0 {
		return true
	}
	if memo[index][c] != -1 {
		return memo[index][c] == 1
	}
	// 1、不将 nums[index] 放入背包
	res := tryPartition(nums, index-1, c)
	// 2、将 nums[index] 放入背包
	if nums[index] <= c {
		res = res || tryPartition(nums, index-1, c-nums[index])
	}
	if res {
		memo[index][c] = 1
	} else {
		memo[index][c] = 0
	}
	return res
}*/

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	n := len(nums)
	c := sum / 2
	// dp[j] 表示从 nums 中取出若干个物品放入背包中，是否能否刚好填满容量为 j 的背包
	dp := make([]bool, c+1)
	for j := 0; j <= c; j++ {
		dp[j] = j == nums[0]
	}
	for i := 1; i < n; i++ {
		for j := c; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[c]
}

func Test() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}

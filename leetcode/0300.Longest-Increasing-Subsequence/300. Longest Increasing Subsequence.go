package solution0300

import "fmt"

/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order
of the remaining elements. For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].



Example 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.


Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4


Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1


Constraints:
1 <= nums.length <= 2500
-10^4 <= nums[i] <= 10^4


Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?
*/

/*func lengthOfLIS(nums []int) int {
	n := len(nums)
	// dp[i] 表示以 nums[i] 结尾的最长上升子序列的元素个数
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = maxOf(dp[i], 1+dp[j])
			}
		}
	}
	res := dp[0]
	for i := 1; i < n; i++ {
		res = maxOf(res, dp[i])
	}
	return res
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}*/

var memo []int

func lengthOfLIS(nums []int) int {
	memo = make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	res := lis(nums, 0)
	for i := 1; i < len(nums); i++ {
		res = maxOf(res, lis(nums, i))
	}
	return res
}

// 以 nums[index] 结尾的最长上升子序列
func lis(nums []int, index int) int {
	if index == 0 {
		return 1
	}
	if memo[index] != -1 {
		return memo[index]
	}
	res := 1
	for i := 0; i < index; i++ {
		if nums[i] < nums[index] {
			res = maxOf(res, 1+lis(nums, i))
		}
	}
	memo[index] = res
	return res
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Test() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7}))
}

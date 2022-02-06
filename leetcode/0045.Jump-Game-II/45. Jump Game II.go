package solution0045

/*
Given an array of non-negative integers nums, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

You can assume that you can always reach the last index.



Example 1:
Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.


Example 2:
Input: nums = [2,3,0,1,4]
Output: 2


Constraints:
1 <= nums.length <= 10^4
0 <= nums[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 执行用时：204 ms, 在所有 Go 提交中击败了 5.57% 的用户
/*func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = n
	}
	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := i-1; j >= 0; j-- {
			if j+nums[j] >= i {
				dp[i] = minOf(dp[i], 1+dp[j])
			}
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
}*/

func jump(nums []int) int {
	n := len(nums)
	maxPos := 0
	end := 0
	steps := 0
	for i := 0; i < n-1; i++ {
		maxPos = maxOf(maxPos, i+nums[i])
		if i == end {
			end = maxPos
			steps++
		}
	}
	return steps
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

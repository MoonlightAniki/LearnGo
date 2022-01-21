package solution0055

import "fmt"

/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.



Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.


Constraints:
1 <= nums.length <= 10^4
0 <= nums[i] <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 记忆化搜索
// 执行用时：412 ms, 在所有 Go 提交中击败了 7.30% 的用户
/*func canJump(nums []int) bool {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = -1
	}
	return jump(nums, 0, dp)
}

func jump(nums []int, start int, dp []int) bool {
	if start >= len(nums)-1 {
		return true
	}
	if dp[start] != -1 {
		return dp[start] == 1
	}
	for i := 1; i <= nums[start]; i++ {
		if jump(nums, start+i, dp) {
			dp[start] = 1
			return true
		}
	}
	dp[start] = 0
	return false
}*/

// 动态规划
// 执行用时：124 ms, 在所有 Go 提交中击败了 13.66% 的用户
/*func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[len(dp)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		for j := nums[i]; j >= 1; j-- {
			if i+j >= len(nums)-1 || dp[i+j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}*/

// 动态规划
// 执行用时：88 ms, 在所有 Go 提交中击败了 19.59% 的用户
/*func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && nums[j] >= i-j {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}*/

// 执行用时：56 ms, 在所有 Go 提交中击败了 84.46% 的用户
func canJump(nums []int) bool {
	for i := len(nums) - 2; i >= 0; i-- {
		// 找到等于0的元素
		if nums[i] != 0 {
			continue
		}

		// 检查i前面有没有元素可以跨过i位置处的0
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > i-j {
				// 从j位置处继续往前寻找等于0的元素
				i = j
				break
			}
		}
		// i前面没有元素可以跨过i位置处的0
		if j == -1 {
			return false
		}
	}
	return true
}

func Test() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

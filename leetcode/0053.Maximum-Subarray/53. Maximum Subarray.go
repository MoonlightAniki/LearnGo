package solution0053

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.



Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:
Input: nums = [1]
Output: 1

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23


Constraints:
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4


Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 贪心算法
func maxSubArray(nums []int) int {
	res := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum+nums[i] < nums[i] {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		res = maxOf(res, sum)
	}
	return res
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

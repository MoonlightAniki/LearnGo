package solution0018

import (
	"fmt"
	"sort"
)

/*
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.



Example 1:
Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

Example 2:
Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]


Constraints:
1 <= nums.length <= 200
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for a := 0; a < len(nums); {
		for b := a + 1; b < len(nums); {
			for c, d := b+1, len(nums)-1; c < d; {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum > target {
					d = prevNumIndex(nums, d)
				} else if sum < target {
					c = nextNumIndex(nums, c)
				} else {
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					c = nextNumIndex(nums, c)
					d = prevNumIndex(nums, d)
				}
			}
			b = nextNumIndex(nums, b)
		}
		a = nextNumIndex(nums, a)
	}
	return res
}

func nextNumIndex(nums []int, i int) int {
	for j := i + 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			return j
		}
	}
	return len(nums)
}

func prevNumIndex(nums []int, i int) int {
	for j := i - 1; j >= 0; j-- {
		if nums[j] != nums[i] {
			return j
		}
	}
	return -1
}

func Test() {
	fmt.Println(
		fourSum([]int{1, 0, -1, 0, -2, 2}, 0),
		fourSum([]int{1, 1, 1, 1, 1, 1}, 4),
	)
}

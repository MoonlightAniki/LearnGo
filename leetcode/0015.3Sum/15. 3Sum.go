package solution0015

import (
	"fmt"
	"sort"
)

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.


Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Example 2:
Input: nums = []
Output: []

Example 3:
Input: nums = [0]
Output: []


Constraints:
0 <= nums.length <= 3000
-10^5 <= nums[i] <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); {
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j = nextNumIndex(nums, j)
			} else if sum > 0 {
				k = prevNumIndex(nums, k)
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j = nextNumIndex(nums, j)
				k = prevNumIndex(nums, k)
			}
		}
		i = nextNumIndex(nums, i)
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
		threeSum([]int{-1, 0, 1, 2, -1, -4}),
		",",
		threeSum([]int{0, 0, 0, 0}),
		",",
		threeSum([]int{1, -1, -1, 0}),
	)
}

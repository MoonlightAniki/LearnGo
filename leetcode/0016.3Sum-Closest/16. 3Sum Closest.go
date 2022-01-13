package solution0016

import (
	"fmt"
	"sort"
)

/*
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.



Example 1:
Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

Example 2:
Input: nums = [0,0,0], target = 1
Output: 0


Constraints:
3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-10^4 <= target <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	res := nums[0] + nums[1] + nums[2]
	a := abs(res - target)
	for i := 0; i < n; {
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum < target {
				j = nextNumIndex(nums, j)
			} else if sum > target {
				k = prevNumIndex(nums, k)
			} else {
				return target
			}
			if abs(sum-target) < a {
				a = abs(sum - target)
				res = sum
			}
		}
		i = nextNumIndex(nums, i)
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
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
	for j := i - 1; i >= 0; i-- {
		if nums[j] != nums[i] {
			return j
		}
	}
	return -1
}

func Test() {
	fmt.Println(
		//threeSumClosest([]int{-1, 2, 1, -4}, 1) == 2,
		threeSumClosest([]int{0, 2, 1, -3}, 1) == 0,
	)
}

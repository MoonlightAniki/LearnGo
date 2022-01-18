package solution0034

import "fmt"

/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.


Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Example 3:
Input: nums = [], target = 0
Output: [-1,-1]


Constraints:
0 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
nums is a non-decreasing array.
-10^9 <= target <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < len(nums) && nums[l] != target {
		l++
	}
	for r >= 0 && nums[r] != target {
		r--
	}
	if l > r {
		return []int{-1, -1}
	} else {
		return []int{l, r}
	}
}
*/

func searchRange(nums []int, target int) []int {
	index := binarySearch(nums, target)
	if index == -1 {
		return []int{-1, -1}
	} else {
		return []int{findStartIndex(nums, index), findEndIndex(nums, index)}
	}
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func findStartIndex(nums []int, index int) int {
	i := index
	for i >= 0 && nums[i] == nums[index] {
		i--
	}
	return i + 1
}

func findEndIndex(nums []int, index int) int {
	i := index
	for i < len(nums) && nums[i] == nums[index] {
		i++
	}
	return i - 1
}

func Test() {
	fmt.Println(
		searchRange([]int{5, 7, 7, 8, 8, 10}, 8),
		searchRange([]int{5, 7, 7, 8, 8, 10}, 6),
	)
}

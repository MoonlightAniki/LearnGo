package solution0033

import "fmt"

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.



Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:
Input: nums = [1], target = 0
Output: -1


Constraints:
1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
All values of nums are unique.
nums is an ascending array that is possibly rotated.
-10^4 <= target <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target < nums[mid] {
			if nums[l] <= nums[r] {
				r = mid - 1
			} else {
				return _search(nums, l, r, target)
			}
		} else if target > nums[mid] {
			if nums[l] <= nums[r] {
				l = mid + 1
			} else {
				return _search(nums, l, r, target)
			}
		} else {
			return mid
		}
	}
	return -1
}

func _search(nums []int, l int, r int, target int) int {
	for i := l; i <= r; i++ {
		if target == nums[i] {
			return i
		}
	}
	return -1
}*/

/*func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] < nums[r] { // nums[mid...r] 有序
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else { // nums[l...mid] 有序
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}*/

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	// 先找到最小值的索引，也就是旋转点的索引
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	rotate := lo
	lo, hi = 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		realMid := (mid + rotate) % len(nums)
		if target == nums[realMid] {
			return realMid
		} else if target < nums[realMid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}

// 二分查找，非递归实现
func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 二分查找，递归实现
/*func binarySearch(nums []int, target int) int {
	return _binarySearch(nums, 0, len(nums)-1, target)
}

func _binarySearch(nums []int, l int, r int, target int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if target < nums[mid] {
		return _binarySearch(nums, l, mid-1, target)
	} else if target > nums[mid] {
		return _binarySearch(nums, mid+1, r, target)
	} else {
		return mid
	}
}*/

func Test() {
	//testBinarySearch()
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	//fmt.Println(search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8))
}

func testBinarySearch() {
	fmt.Println(binarySearch([]int{1, 3}, 2))
	fmt.Println(binarySearch([]int{1, 3}, 0))
	fmt.Println(binarySearch([]int{1, 3}, 4))
	fmt.Println(binarySearch([]int{1, 3}, 3))
}

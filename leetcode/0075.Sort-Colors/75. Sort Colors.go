package solution0075

/*
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.



Example 1:
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Example 2:
Input: nums = [2,0,1]
Output: [0,1,2]


Constraints:
n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.


Follow up: Could you come up with a one-pass algorithm using only constant extra space?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-colors
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 统计排序，时间复杂度 O(n)
/*func sortColors(nums []int) {
	red, white, blue := 0, 0, 0
	for _, num := range nums {
		if num == 0 {
			red++
		} else if num == 1 {
			white++
		} else if num == 2 {
			blue++
		}
	}
	i := 0
	for j := 0; j < red; j++ {
		nums[i] = 0
		i++
	}
	for j := 0; j < white; j++ {
		nums[i] = 1
		i++
	}
	for j := 0; j < blue; j++ {
		nums[i] = 2
		i++
	}
}*/

func sortColors(nums []int) {
	// nums[0,lt] == 0, nums[gt,len(nums)-1] == 2
	lt, gt := -1, len(nums)
	for i := 0; i < gt; {
		if nums[i] == 0 {
			nums[lt+1], nums[i] = nums[i], nums[lt+1]
			lt++
			i++
		} else if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			nums[i], nums[gt-1] = nums[gt-1], nums[i]
			gt--
		}
	}
}

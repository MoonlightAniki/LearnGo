package solution0031

import "fmt"

/*
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such an arrangement is impossible, it must rearrange it to the lowest possible order (i.e., sorted in ascending order).

The replacement must be in place and use only constant extra memory.



Example 1:
Input: nums = [1,2,3]
Output: [1,3,2]

Example 2:
Input: nums = [3,2,1]
Output: [1,2,3]

Example 3:
Input: nums = [1,1,5]
Output: [1,5,1]


Constraints:
1 <= nums.length <= 100
0 <= nums[i] <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func nextPermutation(nums []int) {
	i := len(nums) - 1
	for ; i-1 >= 0; i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}
	// nums 为降序数组，是最大的组合，将数组反转得到最小的组合返回
	if i == 0 {
		reverse(nums, 0, len(nums)-1)
		return
	}

	// nums[i...n-1]降序，倒序找到第一个比 nums[i-1] 大的元素
	j := len(nums) - 1
	for ; j >= i; j-- {
		if nums[j] > nums[i-1] {
			nums[i-1], nums[j] = nums[j], nums[i-1]
			reverse(nums, i, len(nums)-1)
			break
		}
	}
}

func reverse(nums []int, l int, r int) {
	for i, j := l, r; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func Test() {
	nums := []int{1, 2, 3}
	nums = []int{3, 2, 1}
	nums = []int{4, 5, 2, 6, 3, 1} // {4, 5, 3, 6, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

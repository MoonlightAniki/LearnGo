package solution0905

import "fmt"

/*
Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.

Return any array that satisfies this condition.

Example 1:
Input: nums = [3,1,2,4]
Output: [2,4,3,1]
Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

Example 2:
Input: nums = [0]
Output: [0]

Constraints:
1 <= nums.length <= 5000
0 <= nums[i] <= 5000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-array-by-parity
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func sortArrayByParity(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]&1 == 1 && nums[j]&1 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			if nums[i]&1 == 0 {
				i++
			}
			if nums[j]&1 == 1 {
				j--
			}
		}
	}
	return nums
}

func Test() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{0}))
}

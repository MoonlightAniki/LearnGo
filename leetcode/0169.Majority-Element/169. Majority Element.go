package solution0169

import "fmt"

/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

Example 1:
Input: nums = [3,2,3]
Output: 3

Example 2:
Input: nums = [2,2,1,1,1,2,2]
Output: 2


Constraints:
n == nums.length
1 <= n <= 5 * 104
-2^31 <= nums[i] <= 2^31 - 1

Follow-up: Could you solve the problem in linear time and in O(1) space?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/majority-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*func majorityElement(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
		if freq[num] > len(nums)/2 {
			return num
		}
	}
	panic("illegal arguments")
}*/

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		panic("illegal arguments")
	}
	count := 1
	maj := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == maj {
			count++
		} else {
			count--
			if count == 0 && i+1 < len(nums) {
				maj = nums[i+1]
			}
		}
	}
	return maj
}

func Test() {
	fmt.Println(majorityElement([]int{1, 2, 2}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

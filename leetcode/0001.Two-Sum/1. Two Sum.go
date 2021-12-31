package solution0001

import "fmt"

// https://leetcode-cn.com/problems/two-sum/
// 1. Two Sum
/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.

Follow-up:Can you come up with an algorithm that is less than O(n2) time complexity?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func twoSum(nums []int, target int) []int {
	record := make(map[int]int)
	for i, v := range nums {
		other := target - v
		if index, ok := record[other]; ok {
			return []int{index, i}
		}
		record[v] = i
	}
	return nil
}

func Test() {
	fmt.Println(
		twoSum([]int{2, 7, 11, 15}, 9),
		twoSum([]int{3, 2, 4}, 6),
		twoSum([]int{3, 3}, 6),
	)
}

// 执行用时： 4 ms , 在所有 Go 提交中击败了 97.13% 的用户

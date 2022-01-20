package solution0046

import "fmt"

/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.



Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:
Input: nums = [1]
Output: [[1]]


Constraints:
1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	cur_res := make([]int, 0)
	used := make([]bool, len(nums))
	backtrack(nums, &cur_res, &res, used)
	return res
}

func backtrack(nums []int, cur_res *[]int, res *[][]int, used []bool) {
	if len(*cur_res) == len(nums) {
		*res = append(*res, append([]int{}, *cur_res...))
		return
	}
	for i, num := range nums {
		if !used[i] {
			used[i] = true
			*cur_res = append(*cur_res, num)
			backtrack(nums, cur_res, res, used)
			*cur_res = (*cur_res)[:len(*cur_res)-1]
			used[i] = false
		}
	}
}

func Test() {
	fmt.Println(permute([]int{1, 2, 3}))
}

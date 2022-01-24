package solution0078

import "fmt"

/*
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Example 2:

Input: nums = [0]
Output: [[],[0]]


Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*func subsets(nums []int) [][]int {
	cur_res := make([]int, 0)
	res := make([][]int, 0)
	for count := 0; count <= len(nums); count++ {
		backtrack(nums, 0, count, &cur_res, &res)
	}
	return res
}

func backtrack(nums []int, start int, count int, cur_res *[]int, res *[][]int) {
	if len(*cur_res) == count {
		*res = append(*res, append([]int{}, *cur_res...))
		return
	}
	for i := start; i < len(nums); i++ {
		*cur_res = append(*cur_res, nums[i])
		backtrack(nums, i+1, count, cur_res, res)
		*cur_res = (*cur_res)[:len(*cur_res)-1]
	}
}*/

func subsets(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	for state := 0; state < (1 << n); state++ {
		res = append(res, getSubset(nums, state))
	}
	return res
}

func getSubset(nums []int, state int) []int {
	res := make([]int, 0)
	k := 0
	for state > 0 {
		if state&1 == 1 {
			res = append(res, nums[k])
		}
		state >>= 1
		k++
	}
	return res
}

func Test() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{1}))
}

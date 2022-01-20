package solution0039

import "fmt"

/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.



Example 1:
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

Example 2:
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:
Input: candidates = [2], target = 1
Output: []


Constraints:
1 <= candidates.length <= 30
1 <= candidates[i] <= 200
All elements of candidates are distinct.
1 <= target <= 500

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	cur_res := make([]int, 0)
	solve(candidates, 0, target, &cur_res, &res)
	return res
}

func solve(candidates []int, start int, target int, cur_res *[]int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int{}, *cur_res...))
		return
	}
	for i := start; i < len(candidates); i++ {
		if target >= candidates[i] {
			*cur_res = append(*cur_res, candidates[i])
			solve(candidates, i, target-candidates[i], cur_res, res)
			*cur_res = (*cur_res)[:len(*cur_res)-1]
		}
	}
}

func Test() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

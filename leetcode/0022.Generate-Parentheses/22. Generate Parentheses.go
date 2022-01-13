package solution0022

import "fmt"

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:
Input: n = 1
Output: ["()"]

Constraints:
1 <= n <= 8

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		res = append(res, "")
		return res
	}
	if n == 1 {
		res = append(res, "()")
		return res
	}
	for i := 0; i < n; i++ {
		for _, in := range generateParenthesis(i) {
			for _, out := range generateParenthesis(n - 1 - i) {
				res = append(res, fmt.Sprintf("(%s)%s", in, out))
			}
		}
	}
	return res
}

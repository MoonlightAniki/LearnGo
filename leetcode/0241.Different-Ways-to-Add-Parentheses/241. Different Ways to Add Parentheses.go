package solution0241

import "strconv"

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
func diffWaysToCompute(expression string) []int {
	res := make([]int, 0)
	for i, c := range expression {
		if c == '*' || c == '+' || c == '-' {
			res1, res2 := diffWaysToCompute(expression[:i]), diffWaysToCompute(expression[i+1:])
			for _, r1 := range res1 {
				for _, r2 := range res2 {
					if c == '*' {
						res = append(res, r1*r2)
					} else if c == '+' {
						res = append(res, r1+r2)
					} else if c == '-' {
						res = append(res, r1-r2)
					}
				}
			}
		}
	}
	if len(res) == 0 {
		num, _ := strconv.Atoi(expression)
		res = append(res, num)
	}
	return res
}

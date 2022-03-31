package solution0166

import (
	"strconv"
)

// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
func fractionToDecimal(a int, b int) string {
	if a == 0 {
		return "0"
	}
	res := ""
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		res += "-"
	}
	a = abs(a)
	b = abs(b)
	res += strconv.Itoa(a / b)

	if a%b != 0 {
		res += "."
	}

	m := make(map[int]int)
	for {
		a = (a % b) * 10
		// 不存在循环
		if a == 0 {
			return res
		} else {
			if index, ok := m[a]; !ok {
				m[a] = len(res)
				res += strconv.Itoa(a / b)
			} else {
				// 出现了循环
				return res[:index] + "(" + res[index:] + ")"
			}
		}
	}
}

func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}

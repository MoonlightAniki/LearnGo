package solution0032

import "fmt"

/*
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.



Example 1:
Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".

Example 2:
Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".

Example 3:
Input: s = ""
Output: 0


Constraints:
0 <= s.length <= 3 * 10^4
s[i] is '(', or ')'.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 超出时间限制
/*func longestValidParentheses(s string) int {
	length := len(s)
	if length&1 == 1 {
		length--
	}
	bytes := []byte(s)
	for ; length > 0; length -= 2 {
		for i := 0; i < len(bytes); i++ {
			l, r := i, i+length-1
			if r == len(bytes) {
				break
			}
			if isValidParentheses(bytes, l, r) {
				return length
			}
		}
	}
	return 0
}

func isValidParentheses(s []byte, l int, r int) bool {
	stack := make([]byte, 0)
	for i := l; i <= r; i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}*/

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	// 匹配成功的标记为0，无法匹配的标记为1，最终就是要求出连续的0的最大长度
	mark := make([]int, len(s))
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				// 无法匹配的右括号标记为1
				mark[i] = 1
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	for _, i := range stack {
		mark[i] = 1
	}
	res := 0
	length := 0
	for _, m := range mark {
		if m == 1 {
			length = 0
		} else {
			length++
		}
		res = maxOf(res, length)
	}
	return res
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Test() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("(()())"))
}

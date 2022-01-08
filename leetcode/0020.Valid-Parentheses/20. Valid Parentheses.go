package solution0020

import (
	"LearnGo/data_structures/03-Stacks-and-Queues/02-Array-Stack/arraystack"
	"fmt"
)

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Constraints:
1 <= s.length <= 104
s consists of parentheses only '()[]{}'.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := arraystack.NewArrayStack()
	for _, b := range []byte(s) {
		if b == '(' || b == '[' || b == '{' {
			stack.Push(b)
		} else {
			if stack.IsEmpty() {
				return false
			}
			top := stack.Pop().(byte)
			if b == ')' && top != '(' {
				return false
			}
			if b == ']' && top != '[' {
				return false
			}
			if b == '}' && top != '{' {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func Test() {
	fmt.Println(
		isValid("()"),
		isValid("()[]{}"),
		isValid("(]"),
	)
}

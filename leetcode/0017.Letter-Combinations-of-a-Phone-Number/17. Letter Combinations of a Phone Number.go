package solution0017

import "fmt"

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.



Example 1:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:
Input: digits = ""
Output: []

Example 3:
Input: digits = "2"
Output: ["a","b","c"]


Constraints:
0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*func letterCombinations(digits string) []string {
	return helper([]byte(digits), 0)
}

var m = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

// 返回digits[i...n)范围内数字可得到的所有字母组合
func helper(digits []byte, index int) []string {
	res := make([]string, 0)
	if index == len(digits) {
		return res
	}
	if index == len(digits)-1 {
		res = append(res, m[digits[index]]...)
	}
	for _, s := range m[digits[index]] {
		for _, t := range helper(digits, index+1) {
			res = append(res, s+t)
		}
	}
	return res
}*/

var letterMap = []string{
	"",
	"",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

var res = make([]string, 0)

func letterCombinations(digits string) []string {
	res = make([]string, 0)
	if digits == "" {
		return res
	}
	findCombination([]byte(digits), 0, "")
	return res
}

func findCombination(digits []byte, index int, s string) {
	if index == len(digits) {
		res = append(res, s)
		return
	}
	for _, ch := range letterMap[digits[index]-'0'] {
		findCombination(digits, index+1, s+string(ch))
	}
}

func Test() {
	fmt.Println(
		letterCombinations(""),
		letterCombinations("2"),
		letterCombinations("23"),
	)
}

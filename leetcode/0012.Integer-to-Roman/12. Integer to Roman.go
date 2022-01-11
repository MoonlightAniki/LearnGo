package solution0012

import (
	"fmt"
)

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II.
The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead,
the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to
the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral.


Example 1:
Input: num = 3
Output: "III"
Explanation: 3 is represented as 3 ones.

Example 2:
Input: num = 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.

Example 3:
Input: num = 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.


Constraints:
1 <= num <= 3999

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-to-roman
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*var m = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

// 执行用时：12 ms , 在所有 Go 提交中击败了 49.20% 的用户
func intToRoman(num int) string {
	keys := make([]int, len(m))
	i := 0
	for k, _ := range m {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	buffer := bytes.Buffer{}
	for i := len(keys) - 1; i >= 0; i-- {
		key := keys[i]
		value := m[key]
		for num >= key {
			buffer.WriteString(value)
			num -= key
		}
	}
	return buffer.String()
}*/

var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
func intToRoman(num int) string {
	res := ""
	for i, value := range values {
		for num >= value {
			res += symbols[i]
			num -= value
		}
	}
	return res
}

func Test() {
	fmt.Println(
		intToRoman(3) == "III",
		intToRoman(58) == "LVIII",
		intToRoman(1994) == "MCMXCIV",
	)
}

package solution0005

import "fmt"

/*
Given a string s, return the longest palindromic substring in s.

Example 1:
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"

Constraints:
1 <= s.length <= 1000
s consist of only digits and English letters.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 时间复杂度 O(n^3)
/*func longestPalindrome(s string) string {
	bytes := []byte(s)
	for length := len(bytes); length > 0; length-- {
		for l := 0; l+length-1 < len(bytes); l++ {
			r := l + length - 1
			if isPalindrome(bytes, l, r) {
				return string(bytes[l : r+1])
			}
		}
	}
	return ""
}

func isPalindrome(s []byte, l int, r int) bool {
	i, j := l, r
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}*/

// 时间复杂度 O(n^2)
func longestPalindrome(s string) string {
	bytes := []byte(s)
	res := ""
	for i := 0; i < len(bytes); i++ {
		p1 := findPalindrome(bytes, i, i)
		if len(p1) > len(res) {
			res = p1
		}
		if i+1 < len(bytes) {
			p2 := findPalindrome(bytes, i, i+1)
			if len(p2) > len(res) {
				res = p2
			}
		}
	}
	return res
}

func findPalindrome(s []byte, l int, r int) string {
	for l >= 0 && r < len(s) {
		if s[l] == s[r] {
			l--
			r++
		} else {
			return string(s[l+1 : r])
		}
	}
	return string(s[l+1 : r])
}

func Test() {
	fmt.Println(longestPalindrome("babad"))
}

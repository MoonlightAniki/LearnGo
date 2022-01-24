package solution0076

import "fmt"

/*
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character
in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

A substring is a contiguous sequence of characters within the string.



Example 1:
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Example 2:
Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.

Example 3:
Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.


Constraints:
m == s.length
n == t.length
1 <= m, n <= 10^5
s and t consist of uppercase and lowercase English letters.


Follow up: Could you find an algorithm that runs in O(m + n) time?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 超出时间限制
/*func minWindow(s string, t string) string {
	m, n := len(s), len(t)
	if m < n {
		return ""
	}
	tFreq := make(map[rune]int)
	for _, c := range t {
		tFreq[c]++
	}
	var sFreq map[rune]int
	runes := []rune(s)
	res := ""
	for i := 0; i < m; i++ {
		if tFreq[runes[i]] != 0 {
			sFreq = make(map[rune]int)
			for j := i; j < m; j++ {
				sFreq[runes[j]]++
				if contains(sFreq, tFreq) {
					if j-i+1 < len(res) || res == "" {
						res = string(runes[i:(j + 1)])
					}
					break
				}
			}
		}
	}
	return res
}

func contains(m1, m2 map[rune]int) bool {
	if len(m2) > len(m1) {
		return false
	}
	for k, v := range m2 {
		if m1[k] < v {
			return false
		}
	}
	return true
}*/

// 执行用时：20 ms, 在所有 Go 提交中击败了 77.27% 的用户
// 滑动窗口，时间复杂度 O(m+n)
/*func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tFreq := make(map[byte]int)
	for _, b := range []byte(t) {
		tFreq[b]++
	}
	sFreq := make(map[byte]int)
	l, r := 0, -1
	startIndex := -1
	minLength := len(s) + 1
	sCnt := 0
	bytes := []byte(s)
	for l < len(s) {
		if r+1 < len(s) && sCnt < len(t) {
			sFreq[bytes[r+1]]++
			if sFreq[bytes[r+1]] <= tFreq[bytes[r+1]] {
				sCnt++
			}
			r++
		} else {
			if sCnt == len(t) && r-l+1 < minLength {
				minLength = r - l + 1
				startIndex = l
			}
			sFreq[bytes[l]]--
			if sFreq[bytes[l]] < tFreq[bytes[l]] {
				sCnt--
			}
			l++
		}
	}
	if startIndex != -1 {
		return string(bytes[startIndex : startIndex+minLength])
	}
	return ""
}*/

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
/*func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tFreq := make([]int, 256)
	for _, b := range []byte(t) {
		tFreq[b]++
	}

	l, r := 0, -1
	startIndex := -1
	minLength := len(s) + 1
	sCnt := 0
	bytes := []byte(s)
	sFreq := make([]int, 256)
	for l < len(s) {
		if r+1 < len(s) && sCnt < len(t) {
			sFreq[bytes[r+1]]++
			if sFreq[bytes[r+1]] <= tFreq[bytes[r+1]] {
				sCnt++
			}
			r++
		} else {
			if sCnt == len(t) && r-l+1 < minLength {
				minLength = r - l + 1
				startIndex = l
			}
			sFreq[bytes[l]]--
			if sFreq[bytes[l]] < tFreq[bytes[l]] {
				sCnt--
			}
			l++
		}
	}
	if startIndex != -1 {
		return string(bytes[startIndex : startIndex+minLength])
	} else {
		return ""
	}
}*/

// 执行用时：8 ms, 在所有 Go 提交中击败了 91.00% 的用户
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	t_set := make(map[byte]bool)
	tFreq := make([]int, 256)
	for _, b := range []byte(t) {
		t_set[b] = true
		tFreq[b]++
	}

	filtered_s := make([]byte, 0)
	pos := make([]int, 0)
	for index, b := range []byte(s) {
		if t_set[b] {
			filtered_s = append(filtered_s, b)
			pos = append(pos, index)
		}
	}
	l, r := 0, -1
	startIndex := -1
	minLength := len(s) + 1
	sCnt := 0
	sFreq := make([]int, 256)
	for l < len(filtered_s) {
		if r+1 < len(filtered_s) && sCnt < len(t) {
			sFreq[filtered_s[r+1]]++
			if sFreq[filtered_s[r+1]] <= tFreq[filtered_s[r+1]] {
				sCnt++
			}
			r++
		} else {
			if sCnt == len(t) && pos[r]-pos[l]+1 < minLength {
				minLength = pos[r] - pos[l] + 1
				startIndex = pos[l]
			}
			sFreq[filtered_s[l]]--
			if sFreq[filtered_s[l]] < tFreq[filtered_s[l]] {
				sCnt--
			}
			l++
		}
	}
	if startIndex != -1 {
		return string([]byte(s)[startIndex : startIndex+minLength])
	} else {
		return ""
	}
}

func Test() {
	fmt.Println(minWindow("a", "b"))
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd"))
}

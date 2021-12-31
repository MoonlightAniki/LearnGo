package main

import "fmt"

/*
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

3. Longest Substring Without Repeating Characters

Given a string s, find the length of the longest substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
执行用时：12 ms, 在所有 Go 提交中击败了45.97%的用户
*/
//func lengthOfLongestSubstring(s string) int {
//	freq := make(map[byte]int)
//	start := 0
//	res := 0
//	bytes := []byte(s)
//	for i := 0; i < len(bytes); {
//		if freq[bytes[i]] == 0 {
//			freq[bytes[i]]++
//			res = maxOf(res, i-start+1)
//			i++
//		} else {
//			freq[bytes[start]]--
//			start++
//		}
//	}
//	return res
//}

/*
支持中文
执行用时：16 ms, 在所有 Go 提交中击败了19.52%的用户
*/
func lengthOfLongestSubstring(s string) int {
	freq := make(map[rune]int)
	start := 0
	res := 0
	chars := []rune(s)
	for i := 0; i < len(chars); {
		if freq[chars[i]] == 0 {
			freq[chars[i]]++
			res = maxOf(res, i-start+1)
			i++
		} else {
			freq[chars[start]]--
			start++
		}
	}
	return res
}

func main() {
	fmt.Println(
		lengthOfLongestSubstring("abcabcbb"),
		lengthOfLongestSubstring("bbbbb"),
		lengthOfLongestSubstring("pwwkew"),
		lengthOfLongestSubstring("我爱慕课网"),
		lengthOfLongestSubstring("一二三二一"),
	)
}

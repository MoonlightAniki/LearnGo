package solution0003

import "fmt"

/*
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

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lengthOfLongestSubstring(s string) int {
	freq := make(map[byte]int)
	// s[start...i)为无重复字符的子串
	start := 0
	res := 0
	bytes := []byte(s)
	for i := 0; i < len(bytes); {
		if freq[bytes[i]] == 0 {
			res = maxOf(res, i-start+1)
			freq[bytes[i]]++
			i++
		} else {
			freq[bytes[start]]--
			start++
		}
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
	fmt.Println(
		lengthOfLongestSubstring("abcabcbb") == 3,
		lengthOfLongestSubstring("bbbbb") == 1,
		lengthOfLongestSubstring("pwwkew") == 3,
	)
}

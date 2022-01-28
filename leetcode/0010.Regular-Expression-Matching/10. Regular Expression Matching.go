package solution0010

import "fmt"

/*
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).



Example 1:
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".


Example 2:
Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".


Example 3:
Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".


Constraints:
1 <= s.length <= 20
1 <= p.length <= 30
s contains only lowercase English letters.
p contains only lowercase English letters, '.', and '*'.
It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
*/

func isMatch(s string, p string) bool {
	return match([]byte(s), []byte(p), 0, 0)
}

func match(s []byte, p []byte, sIndex int, pIndex int) bool {
	fmt.Printf("s:%v, p:%v, sIndex:%v, pIndex:%v\n", s, p, sIndex, pIndex)
	if sIndex == len(s) {
		return noMoreMatch(p, pIndex)
	}
	if pIndex == len(p) {
		return false
	}
	if pIndex+1 < len(p) && p[pIndex+1] == '*' {
		if s[sIndex] == p[pIndex] || p[pIndex] == '.' {
			return match(s, p, sIndex+1, pIndex) || match(s, p, sIndex, pIndex+2)
		} else {
			return match(s, p, sIndex, pIndex+2)
		}
	}
	if s[sIndex] == p[pIndex] || p[pIndex] == '.' {
		return match(s, p, sIndex+1, pIndex+1)
	}
	return false
}

func noMoreMatch(p []byte, pIndex int) bool {
	var i int
	for i = pIndex; i < len(p); i += 2 {
		if i+1 < len(p) && p[i+1] != '*' {
			return false
		}
	}
	return i == len(p)
}

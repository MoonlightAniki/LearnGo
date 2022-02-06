package solution0097

/*
Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.

An interleaving of two strings s and t is a configuration where they are divided into non-empty substrings such that:

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
Note: a + b is the concatenation of strings a and b.



Example 1:
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true


Example 2:
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
Output: false


Example 3:
Input: s1 = "", s2 = "", s3 = ""
Output: true


Constraints:
0 <= s1.length, s2.length <= 100
0 <= s3.length <= 200
s1, s2, and s3 consist of lowercase English letters.


Follow up: Could you solve it using only O(s2.length) additional memory space?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/interleaving-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return dfs([]byte(s1), []byte(s2), []byte(s3), 0, 0, dp)
}

func dfs(s1 []byte, s2 []byte, s3 []byte, i int, j int, dp [][]int) bool {
	if i == len(s1) && j == len(s2) {
		dp[i][j] = 1
		return true
	}
	if dp[i][j] != -1 {
		return dp[i][j] == 1
	}
	if i == len(s1) {
		if s2[j] == s3[i+j] && dfs(s1, s2, s3, i, j+1, dp) {
			dp[i][j] = 1
			return true
		} else {
			dp[i][j] = 0
			return false
		}
	}
	if j == len(s2) {
		if s1[i] == s3[i+j] && dfs(s1, s2, s3, i+1, j, dp) {
			dp[i][j] = 1
			return true
		} else {
			dp[i][j] = 0
			return false
		}
	}
	if s1[i] == s3[i+j] && dfs(s1, s2, s3, i+1, j, dp) {
		dp[i][j] = 1
		return true
	}
	if s2[j] == s3[i+j] && dfs(s1, s2, s3, i, j+1, dp) {
		dp[i][j] = 1
		return true
	}
	dp[i][j] = 0
	return false
}

package solution0139

import "fmt"

/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.



Example 1:
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

Example 3:
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false


Constraints:
1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s and wordDict[i] consist of only lowercase English letters.
All the strings of wordDict are unique.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func wordBreak(s string, wordDict []string) bool {
	// dp[i]表示 s[i:] 是否能被拆分
	dp := make([]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	return dfs([]byte(s), 0, wordDict, dp)
}

func dfs(bytes []byte, index int, wordDict []string, dp []int) bool {
	if index == len(bytes) {
		return true
	}
	if dp[index] != -1 {
		return dp[index] == 1
	}
	for _, word := range wordDict {
		if index+len(word) <= len(bytes) &&
			string(bytes[index:index+len(word)]) == word &&
			dfs(bytes, index+len(word), wordDict, dp) {
			dp[index] = 1
			return true
		}
	}
	dp[index] = 0
	return false
}

/*func wordBreak(s string, wordDict []string) bool {
	bytes := []byte(s)
	n := len(s)
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	// dp[i] 表示 s[0...i-1] 能否被 wordDict 拆分, dp[n] 即结果
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// s[0...j-1] 能被拆分 && s[j...i-1] 在 wordDict 中
			if dp[j] && wordSet[string(bytes[j:i])] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}*/

func Test() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

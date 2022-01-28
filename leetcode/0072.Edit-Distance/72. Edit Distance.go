package solution0072

import "math"

/*
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:

Insert a character
Delete a character
Replace a character


Example 1:
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')


Example 2:
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')


Constraints:
0 <= word1.length, word2.length <= 500
word1 and word2 consist of lowercase English letters.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 记忆化搜索
// 执行用时：8 ms, 在所有 Go 提交中击败了 42.18% 的用户
/*var memo [][]int

func minDistance(word1 string, word2 string) int {
	memo = make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		memo[i] = make([]int, len(word2))
		for j := 0; j < len(word2); j++ {
			memo[i][j] = -1
		}
	}
	return dfs([]byte(word1), []byte(word2), 0, 0)
}

func dfs(word1 []byte, word2 []byte, index1 int, index2 int) int {
	if index1 == len(word1) {
		return len(word2) - index2
	}
	if index2 == len(word2) {
		return len(word1) - index1
	}
	if memo[index1][index2] != -1 {
		return memo[index1][index2]
	}
	res := math.MaxInt
	if word1[index1] == word2[index2] {
		res = minOf(res, dfs(word1, word2, index1+1, index2+1))
	}
	res = minOf(res, 1+dfs(word1, word2, index1, index2+1))   // word1 插入一个字符
	res = minOf(res, 1+dfs(word1, word2, index1+1, index2))   // word1 删除一个字符
	res = minOf(res, 1+dfs(word1, word2, index1+1, index2+1)) // word1 替换一个字符
	memo[index1][index2] = res
	return res
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}*/

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// dp[i][j] 表示从 word1[0...i) 变成 word2[0...j) 需要最少的操作数
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	bytes1, bytes2 := []byte(word1), []byte(word2)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if bytes1[i-1] == bytes2[j-1] {
				dp[i][j] = minOf(dp[i][j], dp[i-1][j-1])
			}
			dp[i][j] = minOf(dp[i][j], 1+dp[i][j-1])
			dp[i][j] = minOf(dp[i][j], 1+dp[i-1][j])
			dp[i][j] = minOf(dp[i][j], 1+dp[i-1][j-1])
		}
	}
	return dp[m][n]
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

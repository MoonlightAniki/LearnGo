package solution0096

import "fmt"

/*
Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly n nodes of unique values from 1 to n.



Example 1:
Input: n = 3
Output: 5

Example 2:
Input: n = 1
Output: 1


Constraints:
1 <= n <= 19

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		// l, r 分别为左右子树的节点数
		for l := 0; l < i; l++ {
			r := i - 1 - l
			dp[i] += dp[l] * dp[r]
		}
	}
	return dp[n]
}

func Test() {
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(2))
}

package solution0104

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxOf(maxDepth(root.Left), maxDepth(root.Right))
}

package solution0104

import (
	"LearnGo/leetcode/kit"
)

type TreeNode = kit.TreeNode

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 递归解法
/*func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxOf(maxDepth(root.Left), maxDepth(root.Right))
}*/

// BFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	depth := 0
	for len(q) > 0 {
		depth++
		count := len(q)
		for i := 0; i < count; i++ {
			node := q[i]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[count:]
	}
	return depth
}

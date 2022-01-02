package solution0094

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

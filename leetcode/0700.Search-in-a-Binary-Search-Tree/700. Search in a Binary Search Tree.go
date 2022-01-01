package solution0700

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	switch {
	case val == root.Val:
		return root
	case val < root.Val:
		return searchBST(root.Left, val)
	default:
		return searchBST(root.Right, val)
	}
}

package solution0872

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return equals(getLeaves(root1), getLeaves(root2))
}

func equals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func getLeaves(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	if root.Left != nil {
		res = append(res, getLeaves(root.Left)...)
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, root.Val)
	}
	if root.Right != nil {
		res = append(res, getLeaves(root.Right)...)
	}
	return res
}

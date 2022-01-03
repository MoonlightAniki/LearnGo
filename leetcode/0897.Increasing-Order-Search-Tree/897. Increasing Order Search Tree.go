package solution0897

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	values := inorderTraverse(root)
	newRoot := &TreeNode{Val: values[0]}
	cur := newRoot
	for i := 1; i < len(values); i++ {
		cur.Right = &TreeNode{Val: values[i]}
		cur = cur.Right
	}
	return newRoot
}

func inorderTraverse(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, inorderTraverse(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraverse(root.Right)...)
	return res
}

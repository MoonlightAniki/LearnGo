package solution0094

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

// 递归解法
/*func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}*/

// 迭代解法
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left == nil && node.Right == nil {
			res = append(res, node.Val)
		} else {
			left, right := node.Left, node.Right
			if right != nil {
				stack = append(stack, right)
			}
			node.Left = nil
			node.Right = nil
			stack = append(stack, node)
			if left != nil {
				stack = append(stack, left)
			}
		}
	}
	return res
}

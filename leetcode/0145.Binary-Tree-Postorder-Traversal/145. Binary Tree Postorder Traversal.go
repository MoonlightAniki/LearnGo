package solution0145

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

// 递归解法
/*func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}*/

// 迭代解法
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left, right := node.Left, node.Right
		if left == nil && right == nil {
			res = append(res, node.Val)
		} else {
			node.Left = nil
			node.Right = nil
			stack = append(stack, node)
			if right != nil {
				stack = append(stack, right)
			}
			if left != nil {
				stack = append(stack, left)
			}
		}
	}
	return res
}

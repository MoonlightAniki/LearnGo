package solution0144

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

// 递归解法
/*func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}*/

// 迭代解法
func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left, right := node.Left, node.Right
		if left == nil && right == nil {
			res = append(res, node.Val)
		} else {
			if right != nil {
				stack = append(stack, right)
			}
			if left != nil {
				stack = append(stack, left)
			}
			node.Left = nil
			node.Right = nil
			stack = append(stack, node)
		}
	}
	return res
}

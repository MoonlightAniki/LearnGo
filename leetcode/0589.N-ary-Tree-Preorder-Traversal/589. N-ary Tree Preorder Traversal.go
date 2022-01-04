package solution0589

import "LearnGo/leetcode/kit"

type Node = kit.Node

// 递归解法
/*func preorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	for _, child := range root.Children {
		res = append(res, preorder(child)...)
	}
	return res
}*/

// 迭代解法
func preorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*Node
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(node.Children) == 0 {
			res = append(res, node.Val)
		} else {
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
			node.Children = nil
			stack = append(stack, node)
		}
	}
	return res
}

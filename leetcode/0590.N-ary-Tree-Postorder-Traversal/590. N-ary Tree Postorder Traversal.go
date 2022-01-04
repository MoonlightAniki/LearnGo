package solution0590

import "LearnGo/leetcode/kit"

type Node = kit.Node

// 递归解法
/*func postorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	for _, node := range root.Children {
		res = append(res, postorder(node)...)
	}
	res = append(res, root.Val)
	return res
}*/

// 迭代解法
func postorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(node.Children) == 0 {
			res = append(res, node.Val)
		} else {
			children := node.Children
			node.Children = nil
			stack = append(stack, node)
			for i := len(children) - 1; i >= 0; i-- {
				stack = append(stack, children[i])
			}
		}
	}
	return res
}

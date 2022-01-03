package solution0589

import "LearnGo/leetcode/kit"

type Node = kit.Node

func preorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	for _, child := range root.Children {
		res = append(res, preorder(child)...)
	}
	return res
}

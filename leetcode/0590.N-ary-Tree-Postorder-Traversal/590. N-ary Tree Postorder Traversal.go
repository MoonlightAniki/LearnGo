package solution0590

import "LearnGo/leetcode/kit"

type Node = kit.Node

func postorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	for _, node := range root.Children {
		res = append(res, postorder(node)...)
	}
	res = append(res, root.Val)
	return res
}

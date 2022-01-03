package solution0559

import "LearnGo/leetcode/kit"

type Node = kit.Node

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	childMaxDepth := 0
	for _, child := range root.Children {
		childMaxDepth = maxOf(childMaxDepth, maxDepth(child))
	}
	return childMaxDepth + 1
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

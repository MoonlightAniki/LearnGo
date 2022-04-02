package solution0662

import (
	"LearnGo/leetcode/kit"
)

type TreeNode = kit.TreeNode

// 执行用时：4 ms , 在所有 Go 提交中击败了 85.24% 的用户
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := []*TreeNode{root}
	indexes := []int{0}
	res := 0
	for len(nodes) > 0 {
		res = maxOf(res, indexes[len(indexes)-1]-indexes[0]+1)

		length := len(nodes)
		for i := 0; i < length; i++ {
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
				indexes = append(indexes, indexes[i]*2+1)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
				indexes = append(indexes, indexes[i]*2+2)
			}
		}
		nodes = nodes[length:]
		indexes = indexes[length:]
	}
	return res
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

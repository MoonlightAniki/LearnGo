package solution0637

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		sum := 0
		count := len(queue)
		for i := 0; i < count; i++ {
			node := queue[i]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, float64(sum)/float64(count))
		queue = queue[count:]
	}
	return res
}

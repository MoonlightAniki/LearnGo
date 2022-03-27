package solution0222

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

// 执行用时： 16 ms , 在所有 Go 提交中击败了 58.47% 的用户
func countNodes(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	count := 0
	for len(queue) > 0 {
		length := len(queue)
		count += length
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
	}
	return count
}

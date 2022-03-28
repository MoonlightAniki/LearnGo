package solution0230

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

// 执行用时： 8 ms , 在所有 Go 提交中击败了 85.61% 的用户
func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	rank := 0
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Left == nil && top.Right == nil {
			rank++
			if rank == k {
				return top.Val
			}
		} else {
			leftNode, rightNode := top.Left, top.Right
			top.Left = nil
			top.Right = nil
			if rightNode != nil {
				stack = append(stack, rightNode)
			}
			stack = append(stack, top)
			if leftNode != nil {
				stack = append(stack, leftNode)
			}
		}
	}
	panic("illegal arguments")
}

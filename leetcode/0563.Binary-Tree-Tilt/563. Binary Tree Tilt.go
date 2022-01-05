package solution0563

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

// 先计算出每个节点的和，再层序遍历计算出每个节点左右子节点和的差值
/*func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	root = sum(root)
	tilt := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		left := 0
		if node.Left != nil {
			left = node.Left.Val
			queue = append(queue, node.Left)
		}
		right := 0
		if node.Right != nil {
			right = node.Right.Val
			queue = append(queue, node.Right)
		}
		tilt += abs(left - right)
	}
	return tilt
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func sum(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil {
		root.Left = sum(root.Left)
		root.Val += root.Left.Val
	}
	if root.Right != nil {
		root.Right = sum(root.Right)
		root.Val += root.Right.Val
	}
	return root
}*/

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return abs(sum(root.Left)-sum(root.Right)) + findTilt(root.Left) + findTilt(root.Right)
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + sum(root.Left) + sum(root.Right)
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

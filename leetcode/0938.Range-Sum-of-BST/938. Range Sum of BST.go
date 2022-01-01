// https://leetcode-cn.com/problems/range-sum-of-bst/

package solution0938

import (
	"LearnGo/leetcode/kit"
	"fmt"
)

type TreeNode = kit.TreeNode

// V1
//func rangeSumBST(root *TreeNode, low int, high int) int {
//	sum := 0
//	inOrderTraverse(root, low, high, &sum)
//	return sum
//}
//
//func inOrderTraverse(node *TreeNode, low int, high int, sum *int) {
//	if node == nil {
//		return
//	}
//	if node.Val >= low {
//		inOrderTraverse(node.Left, low, high, sum)
//	}
//	if node.Val >= low && node.Val <= high {
//		*sum += node.Val
//	}
//	if node.Val <= high {
//		inOrderTraverse(node.Right, low, high, sum)
//	}
//}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	sum := 0
	switch {
	case root.Val < low:
		sum += rangeSumBST(root.Right, low, high)
	case root.Val > high:
		sum += rangeSumBST(root.Left, low, high)
	default:
		sum += rangeSumBST(root.Left, low, high)
		sum += rangeSumBST(root.Right, low, high)
		sum += root.Val
	}
	return sum
}

func Test() {
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 18}
	fmt.Println(rangeSumBST(root, 7, 15) == 32)
}

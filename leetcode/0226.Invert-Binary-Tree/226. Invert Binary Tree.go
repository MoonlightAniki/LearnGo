// https://leetcode-cn.com/problems/invert-binary-tree/

package solution0226

import "LearnGo/leetcode/kit"

type TreeNode = kit.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	leftNode, rightNode := root.Left, root.Right
	root.Left = invertTree(rightNode)
	root.Right = invertTree(leftNode)

	return root
}

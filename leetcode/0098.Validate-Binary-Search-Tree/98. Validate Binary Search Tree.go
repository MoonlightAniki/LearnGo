package solution0098

import (
	"LearnGo/leetcode/kit"
)

/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:
Input: root = [2,1,3]
Output: true

Example 2:
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.


Constraints:
The number of nodes in the tree is in the range [1, 104].
-2^31 <= Node.val <= 2^31 - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode = kit.TreeNode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 中序遍历
	var stack = []*TreeNode{root}
	var values []int
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Left == nil && top.Right == nil {
			values = append(values, top.Val)
		} else {
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
			stack = append(stack, &TreeNode{Val: top.Val})
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
		}
	}
	// 判断values是否升序
	for i := 0; i+1 < len(values); i++ {
		if values[i] >= values[i+1] {
			return false
		}
	}
	return true
}

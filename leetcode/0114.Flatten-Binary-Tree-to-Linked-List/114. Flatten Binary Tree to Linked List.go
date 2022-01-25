package solution0114

import "LearnGo/leetcode/kit"

/*
Given the root of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
The "linked list" should be in the same order as a pre-order traversal of the binary tree.


Example 1:
Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [0]
Output: [0]


Constraints:
The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100


Follow up: Can you flatten the tree in-place (with O(1) extra space)?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode = kit.TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	values := make([]int, 0)
	preorder(root, &values)
	newRoot := &TreeNode{Val: values[0]}
	cur := newRoot
	for i := 1; i < len(values); i++ {
		cur.Right = &TreeNode{Val: values[i]}
		cur = cur.Right
	}
	*root = *newRoot
}

func preorder(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	*values = append(*values, root.Val)
	preorder(root.Left, values)
	preorder(root.Right, values)
}

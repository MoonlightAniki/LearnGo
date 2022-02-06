package solution0099

import (
	"LearnGo/leetcode/kit"
	"sort"
)

/*
You are given the root of a binary search tree (BST), where the values of exactly two nodes of the tree were swapped by
mistake. Recover the tree without changing its structure.



Example 1:
Input: root = [1,3,null,null,2]
Output: [3,1,null,null,2]
Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3 makes the BST valid.


Example 2:
Input: root = [3,1,4,null,null,2]
Output: [2,1,4,null,null,3]
Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2 and 3 makes the BST valid.


Constraints:
The number of nodes in the tree is in the range [2, 1000].
-2^31 <= Node.val <= 2^31 - 1


Follow up: A solution using O(n) space is pretty straight-forward. Could you devise a constant O(1) space solution?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/recover-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode = kit.TreeNode

func recoverTree(root *TreeNode) {
	values := make([]int, 0)
	inorder(root, &values)
	sort.Ints(values)
	inorderRecover(root, &values)
}

func inorder(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, values)
	*values = append(*values, root.Val)
	inorder(root.Right, values)
}

func inorderRecover(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	inorderRecover(root.Left, values)
	root.Val = (*values)[0]
	*values = (*values)[1:]
	inorderRecover(root.Right, values)
}

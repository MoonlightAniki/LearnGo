package solution0105

import (
	"LearnGo/leetcode/kit"
	"fmt"
)

/*
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is
the inorder traversal of the same tree, construct and return the binary tree.


Example 1:
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Example 2:
Input: preorder = [-1], inorder = [-1]
Output: [-1]


Constraints:
1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode = kit.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(
	preorder []int, preorderL int, preorderR int,
	inorder []int, inorderL int, inorderR int,
) *TreeNode {
	if preorderL > preorderR {
		return nil
	}
	if preorderL == preorderR {
		return &TreeNode{Val: preorder[preorderL]}
	}
	rootValue := preorder[preorderL]
	// 左子树节点数量
	leftCount := 0
	for i := inorderL; i <= inorderR && inorder[i] != rootValue; i++ {
		leftCount++
	}
	root := &TreeNode{Val: rootValue}
	leftPreorderL := preorderL + 1
	leftPreorderR := leftPreorderL + leftCount - 1
	rightPreorderL := leftPreorderR + 1
	rightPreorderR := preorderR

	leftInorderL := inorderL
	leftInorderR := leftInorderL + leftCount - 1
	rightInorderL := leftInorderR + 1 + 1
	rightInorderR := inorderR

	root.Left = build(preorder, leftPreorderL, leftPreorderR, inorder, leftInorderL, leftInorderR)
	root.Right = build(preorder, rightPreorderL, rightPreorderR, inorder, rightInorderL, rightInorderR)
	return root
}

func Test() {
	fmt.Println(buildTree([]int{1, 2}, []int{2, 1}))
}

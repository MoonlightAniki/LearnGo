package solution0337

import "LearnGo/leetcode/kit"

/*
The thief has found himself a new place for his thievery again. There is only one entrance to this area, called root.

Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that all houses
in this place form a binary tree. It will automatically contact the police if two directly-linked houses were broken into
on the same night.

Given the root of the binary tree, return the maximum amount of money the thief can rob without alerting the police.



Example 1:
Input: root = [3,2,3,null,3,null,1]
Output: 7
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.


Example 2:
Input: root = [3,4,5,1,3,null,1]
Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.


Constraints:
The number of nodes in the tree is in the range [1, 10^4].
0 <= Node.val <= 10^4
*/

type TreeNode = kit.TreeNode

/*var memo map[*TreeNode]int

func rob(root *TreeNode) int {
	memo = make(map[*TreeNode]int)
	return tryRob(root)
}

func tryRob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if v, ok := memo[root]; ok {
		return v
	}
	// 不偷取 root 节点的财物
	a := tryRob(root.Left) + tryRob(root.Right)

	// 偷取 root 节点的财物
	b := root.Val
	if root.Left != nil {
		b += tryRob(root.Left.Left) + tryRob(root.Left.Right)
	}
	if root.Right != nil {
		b += tryRob(root.Right.Left) + tryRob(root.Right.Right)
	}
	memo[root] = maxOf(a, b)
	return maxOf(a, b)
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}*/

func rob(root *TreeNode) int {
	include, _ := tryRob(root)
	return include
}

func tryRob(root *TreeNode) (include, exclude int) {
	if root == nil {
		return 0, 0
	}
	includeL, excludeL := tryRob(root.Left)
	includeR, excludeR := tryRob(root.Right)
	exclude = includeL + includeR
	include = maxOf(exclude, root.Val+excludeL+excludeR)
	return include, exclude
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

package solution0237

import "LearnGo/leetcode/kit"

type ListNode = kit.ListNode

// https://leetcode-cn.com/problems/delete-node-in-a-linked-list/

// node 是要删除的节点，而不是头结点
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	// 因为无法获取到要删除节点node的上一个节点
	// 这里将node.Next的值赋值给node，再将node.Next删除，最终得到的链表跟删除node是一样的
	delNode := node.Next
	node.Val = delNode.Val
	node.Next = delNode.Next
	delNode.Next = nil
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户

package solution0206

import (
	"LearnGo/leetcode/kit"
	"fmt"
)

type ListNode = kit.ListNode

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := head
	for cur != nil {
		node := cur
		cur = cur.Next
		node.Next = dummy.Next
		dummy.Next = node
	}
	return dummy.Next
}

func Test() {
	head := kit.CreateListNode([]int{1, 2, 3, 4, 5})
	fmt.Println(head.String())
	head = reverseList(head)
	fmt.Println(head.String())
}

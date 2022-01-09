package solution0002

import (
	"LearnGo/leetcode/kit"
	"fmt"
)

type ListNode = kit.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		a := 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		b := 0
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + carry
		carry = sum / 10
		node := &ListNode{Val: sum % 10}
		cur.Next = node
		cur = cur.Next
	}
	return dummyHead.Next
}

func Test() {
	fmt.Println(
		addTwoNumbers(kit.CreateListNode([]int{2, 4, 3}), kit.CreateListNode([]int{5, 6, 4})),
		"\n",
		addTwoNumbers(kit.CreateListNode([]int{0}), kit.CreateListNode([]int{0})),
		"\n",
		addTwoNumbers(kit.CreateListNode([]int{9, 9, 9, 9, 9, 9, 9}), kit.CreateListNode([]int{9, 9, 9, 9})),
	)
}

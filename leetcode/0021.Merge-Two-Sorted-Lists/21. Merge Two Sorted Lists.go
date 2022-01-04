package solution0021

import (
	"LearnGo/leetcode/kit"
	"fmt"
)

type ListNode = kit.ListNode

// 非递归实现
/*func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummy := &ListNode{}
	cur := dummy
	cur1 := list1
	cur2 := list2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next

			cur = cur.Next
			cur.Next = nil
		} else {
			cur.Next = cur2
			cur2 = cur2.Next

			cur = cur.Next
			cur.Next = nil
		}
	}
	if cur1 != nil {
		cur.Next = cur1
	}
	if cur2 != nil {
		cur.Next = cur2
	}
	return dummy.Next
}*/

// 递归实现
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func Test() {
	fmt.Println(mergeTwoLists(kit.CreateListNode([]int{1, 2, 4}), kit.CreateListNode([]int{1, 3, 4})))
	fmt.Println(mergeTwoLists(kit.CreateListNode([]int{1, 2}), kit.CreateListNode([]int{1, 3, 4, 5})))
}

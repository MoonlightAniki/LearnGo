package kit

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return dummy.Next
}

func (head *ListNode) String() string {
	res := ""
	cur := head
	for cur != nil {
		res += strconv.Itoa(cur.Val) + " -> "
		cur = cur.Next
	}
	res += "nil"
	return res
}

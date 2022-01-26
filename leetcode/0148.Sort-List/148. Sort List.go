package solution0148

import (
	"LearnGo/leetcode/kit"
	"fmt"
)

/*
Given the head of a linked list, return the list after sorting it in ascending order.



Example 1:
Input: head = [4,2,1,3]
Output: [1,2,3,4]


Example 2:
Input: head = [-1,5,3,4,0]
Output: [-1,0,3,4,5]


Example 3:
Input: head = []
Output: []


Constraints:
The number of nodes in the list is in the range [0, 5 * 10^4].
-10^5 <= Node.val <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode = kit.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	leftHead := head
	rightHead := slow.Next
	slow.Next = nil

	leftHead = sortList(leftHead)
	rightHead = sortList(rightHead)

	dummyHead := &ListNode{}
	prev := dummyHead
	for leftHead != nil || rightHead != nil {
		if leftHead == nil {
			prev.Next = rightHead
			break
		}
		if rightHead == nil {
			prev.Next = leftHead
			break
		}
		if leftHead.Val < rightHead.Val {
			node := leftHead
			leftHead = leftHead.Next
			node.Next = nil
			prev.Next = node
			prev = prev.Next
		} else {
			node := rightHead
			rightHead = rightHead.Next
			node.Next = nil
			prev.Next = node
			prev = prev.Next
		}
	}
	ret := dummyHead.Next
	dummyHead.Next = nil
	return ret
}

func Test() {
	fmt.Println(sortList(kit.CreateListNode([]int{4, 2, 1, 3})))
}

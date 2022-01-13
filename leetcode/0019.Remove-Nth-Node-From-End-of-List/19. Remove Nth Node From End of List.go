package solution0019

import (
	"LearnGo/leetcode/kit"
	"fmt"
)

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.


Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]


Constraints:
The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz


Follow up: Could you do this in one pass?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode = kit.ListNode

/*func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := 0
	for cur := head; cur != nil; cur = cur.Next {
		size++
	}
	index := size - n
	if index == 0 {
		return head.Next
	} else {
		prev := head
		for i := 1; i < index; i++ {
			prev = prev.Next
		}
		delNode := prev.Next
		prev.Next = delNode.Next
		delNode.Next = nil
		return head
	}
}*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	fast := dummyHead
	slow := dummyHead
	for i := 1; i <= n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	delNode := slow.Next
	slow.Next = delNode.Next
	delNode.Next = nil
	return dummyHead.Next
}

func Test() {
	fmt.Println(removeNthFromEnd(kit.CreateListNode([]int{1, 2, 3, 4, 5}), 2))
	fmt.Println(removeNthFromEnd(kit.CreateListNode([]int{1, 2, 3, 4, 5}), 5))
	fmt.Println(removeNthFromEnd(kit.CreateListNode([]int{1, 2, 3, 4, 5}), 1))
}

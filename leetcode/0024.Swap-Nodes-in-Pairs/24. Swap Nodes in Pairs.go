package solution0024

import (
	"LearnGo/leetcode/kit"
	"fmt"
)

/*
Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)



Example 1:
Input: head = [1,2,3,4]
Output: [2,1,4,3]

Example 2:
Input: head = []
Output: []

Example 3:
Input: head = [1]
Output: [1]


Constraints:
The number of nodes in the list is in the range [0, 100].
0 <= Node.val <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode = kit.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nn := head.Next.Next
	head.Next.Next = nil
	n := head.Next
	head.Next = nil
	n.Next = head
	n.Next.Next = swapPairs(nn)
	return n
}

func Test() {
	fmt.Println(swapPairs(kit.CreateListNode([]int{1, 2, 3, 4})))
}

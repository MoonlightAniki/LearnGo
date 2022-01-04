package solution0876

import (
	"LearnGo/leetcode/kit"
	"fmt"
)

/*
Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.


Example 1:
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.

Constraints:
The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/middle-of-the-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode = kit.ListNode

func middleNode(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow.Next
}

func Test() {
	fmt.Println(middleNode(kit.CreateListNode([]int{1, 2, 3, 4, 5})))
	fmt.Println(middleNode(kit.CreateListNode([]int{1, 2, 3, 4, 5, 6})))
}

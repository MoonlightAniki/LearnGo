package solution0023

import (
	"LearnGo/leetcode/kit"
	"container/heap"
	"fmt"
)

/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.



Example 1:
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

Example 2:
Input: lists = []
Output: []

Example 3:
Input: lists = [[]]
Output: []


Constraints:
k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] is sorted in ascending order.
The sum of lists[i].length won't exceed 10^4.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode = kit.ListNode

// 执行用时：232 ms , 在所有 Go 提交中击败了 9.33% 的用户
// 时间复杂度 O(K*n)，n 为节点总数
/*func mergeKLists(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{}
	prev := dummyHead
	for {
		index := minIndex(lists)
		if index == -1 {
			break
		}
		list := lists[index]
		node := list
		list = list.Next
		node.Next = nil
		lists[index] = list

		prev.Next = node
		prev = prev.Next
	}
	return dummyHead.Next
}

func minIndex(lists []*ListNode) int {
	index := -1
	var cur *ListNode
	for i, list := range lists {
		if cur == nil && list != nil {
			index = i
			cur = list
		} else if cur != nil && list != nil && list.Val < cur.Val {
			index = i
			cur = list
		}
	}
	return index
}*/

// 执行用时： 12 ms , 在所有 Go 提交中击败了 54.55% 的用户
// 时间复杂度 O(nlogk) n 为节点总数
func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &minHeap{}
	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}
	dummyHead := &ListNode{}
	prev := dummyHead
	for minHeap.Len() > 0 {
		min := heap.Pop(minHeap).(*ListNode)
		if minHeap.Len() == 0 {
			prev.Next = min
			break
		}
		node := min
		min = min.Next
		node.Next = nil
		if min != nil {
			heap.Push(minHeap, min)
		}

		prev.Next = node
		prev = prev.Next
	}
	return dummyHead.Next
}

type minHeap []*ListNode

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}

func (m *minHeap) Pop() interface{} {
	n := len(*m)
	element := (*m)[n-1]
	*m = (*m)[:n-1]
	return element
}

func Test() {
	fmt.Println(mergeKLists([]*ListNode{
		kit.CreateListNode([]int{1, 4, 5}),
		kit.CreateListNode([]int{1, 3, 4}),
		kit.CreateListNode([]int{2, 6}),
	}))
	fmt.Println(mergeKLists(nil))
	fmt.Println(mergeKLists(make([]*ListNode, 0)))
}

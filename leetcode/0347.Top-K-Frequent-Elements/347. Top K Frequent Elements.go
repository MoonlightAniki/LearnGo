package solution0347

import (
	"container/heap"
	"fmt"
)

/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.


Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]


Constraints:
1 <= nums.length <= 10^5
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.


Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/top-k-frequent-elements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 桶排序
// 执行用时：12 ms, 在所有 Go 提交中击败了 92.89% 的用户
/*func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	bucket := make([][]int, len(nums)+1)
	for num, count := range freq {
		bucket[count] = append(bucket[count], num)
	}
	var res []int
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) > 0 {
			for j := 0; j < len(bucket[i]); j++ {
				res = append(res, bucket[i][j])
				if len(res) == k {
					return res
				}
			}
		}
	}
	return res
}*/

type item struct {
	key   int
	count int
}

// 定义别名 并 实现 heap.Interface 实现小顶堆
type priorityQueue []*item

func (pq *priorityQueue) Len() int {
	return len(*pq)
}

func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i].count < (*pq)[j].count
}

func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*item))
}

func (pq *priorityQueue) Pop() interface{} {
	n := len(*pq)
	_item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return _item
}

// 优先队列
// 执行用时：12 ms, 在所有 Go 提交中击败了 92.89% 的用户
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	pq := &priorityQueue{}
	for num, count := range freq {
		heap.Push(pq, &item{num, count})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	var res []int
	for pq.Len() > 0 {
		res = append(res, heap.Pop(pq).(*item).key)
	}
	return res
}

func Test() {
	fmt.Println(
		topKFrequent([]int{4, 1, 1, 1, 1, 2, 2, 2, 3, 3, 5, 5, 5, 5}, 2),
		topKFrequent([]int{1}, 1),
	)
}

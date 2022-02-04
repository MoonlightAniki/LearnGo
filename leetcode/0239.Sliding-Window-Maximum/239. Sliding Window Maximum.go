package solution0239

import "container/heap"

/*
You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the
array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.



Example 1:
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

left  [1, 3, 3, -3, 5, 5, 6, 7]
right [3, 3, -1, 5, 5, 3, 7, 7]

Example 2:
Input: nums = [1], k = 1
Output: [1]


Constraints:
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 440 ms, 在所有 Go 提交中击败了 5.33% 的用户
/*type node struct {
	start int
	end   int
	value int
	left  *node
	right *node
}

type segmentTree struct {
	root *node
	data []int
}

func newSegmentTree(data []int) *segmentTree {
	tree := segmentTree{}
	tree.root = buildSegmentTree(data, 0, len(data)-1)
	tree.data = data
	return &tree
}

func (tree *segmentTree) query(queryL int, queryR int) int {
	return tree.__query(tree.root, queryL, queryR)
}

func (tree *segmentTree) __query(n *node, queryL int, queryR int) int {
	if queryL > queryR || queryL < n.start || queryR > n.end {
		panic("Illegal arguments!")
	}
	if queryL == n.start && queryR == n.end {
		return n.value
	}
	mid := n.start + (n.end-n.start)/2
	if queryR <= mid {
		return tree.__query(n.left, queryL, queryR)
	} else if queryL >= mid+1 {
		return tree.__query(n.right, queryL, queryR)
	} else {
		leftRes := tree.__query(n.left, queryL, mid)
		rightRes := tree.__query(n.right, mid+1, queryR)
		return maxOf(leftRes, rightRes)
	}
}

func buildSegmentTree(data []int, start int, end int) *node {
	if start == end {
		return &node{
			start: start,
			end:   end,
			value: data[start],
		}
	}
	mid := start + (end-start)/2
	left := buildSegmentTree(data, start, mid)
	right := buildSegmentTree(data, mid+1, end)
	value := maxOf(left.value, right.value)
	return &node{
		start: start,
		end:   end,
		value: value,
		left:  left,
		right: right,
	}
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	tree := newSegmentTree(nums)
	var res []int
	for i := 0; i+k-1 < len(nums); i++ {
		res = append(res, tree.query(i, i+k-1))
	}
	return res
}*/

// 执行用时：260 ms, 在所有 Go 提交中击败了 24.40% 的用户
type pair struct {
	num   int
	index int
}

type maxHeap struct {
	data []pair
}

func (m maxHeap) Len() int {
	return len(m.data)
}

func (m maxHeap) Less(i, j int) bool {
	if m.data[i].num == m.data[j].num {
		return m.data[i].index > m.data[j].index
	} else {
		return m.data[i].num > m.data[j].num
	}
}

func (m maxHeap) Swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *maxHeap) Push(x interface{}) {
	m.data = append(m.data, x.(pair))
}

func (m *maxHeap) Pop() interface{} {
	res := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	return res
}

func (m maxHeap) Peek() pair {
	return m.data[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	m := &maxHeap{data: make([]pair, 0)}
	for i := 0; i < k; i++ {
		m.data = append(m.data, pair{nums[i], i})
	}
	heap.Init(m)
	var res []int
	res = append(res, m.Peek().num)
	for i := k; i < len(nums); i++ {
		heap.Push(m, pair{nums[i], i})
		for m.Peek().index <= i-k {
			heap.Pop(m)
		}
		res = append(res, m.Peek().num)
	}
	return res
}

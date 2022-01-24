package solution0084

import "fmt"

/*
Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.



Example 1:
Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.


Example 2:
Input: heights = [2,4]
Output: 4


Constraints:
1 <= heights.length <= 10^5
0 <= heights[i] <= 10^4


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-rectangle-in-histogram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// region 线段树
type node struct {
	rangeL int
	rangeR int
	value  int
	left   *node
	right  *node
}

type segmentTree struct {
	root *node
	data []int
}

func newSegmentTree(data []int) *segmentTree {
	tree := &segmentTree{data: data}
	if len(data) > 0 {
		tree.root = buildSegmentTree(data, 0, len(data)-1)
	}
	return tree
}

func (tree *segmentTree) query(queryL int, queryR int) int {
	return tree.__query(tree.root, queryL, queryR)
}

func (tree *segmentTree) __query(root *node, queryL int, queryR int) int {
	if queryL == root.rangeL && queryR == root.rangeR {
		return root.value
	}
	mid := root.rangeL + (root.rangeR-root.rangeL)/2
	if queryR <= mid {
		return tree.__query(root.left, queryL, queryR)
	} else if queryL >= mid+1 {
		return tree.__query(root.right, queryL, queryR)
	} else {
		resL := tree.__query(root.left, queryL, mid)
		resR := tree.__query(root.right, mid+1, queryR)
		if tree.data[resL] < tree.data[resR] {
			return resL
		} else {
			return resR
		}
	}
}

func buildSegmentTree(data []int, l int, r int) *node {
	n := &node{rangeL: l, rangeR: r}
	if l == r {
		n.value = l
		return n
	}
	mid := l + (r-l)/2
	n.left = buildSegmentTree(data, l, mid)
	n.right = buildSegmentTree(data, mid+1, r)
	if data[n.left.value] < data[n.right.value] {
		n.value = n.left.value
	} else {
		n.value = n.right.value
	}
	return n
}

// endregion

// 执行用时：424 ms , 在所有 Go 提交中击败了 5.25% 的用户
// 线段树 + 分治, 时间复杂度 O(nlogn)
/*func largestRectangleArea(heights []int) int {
	tree := newSegmentTree(heights)
	return largest(heights, 0, len(heights)-1, tree)
}

func largest(heights []int, l int, r int, tree *segmentTree) int {
	if l > r {
		return 0
	}
	if l == r {
		return heights[l]
	}
	minIndex := tree.query(l, r)
	res := (r - l + 1) * heights[minIndex]
	res = maxOf(res, largest(heights, l, minIndex-1, tree))
	res = maxOf(res, largest(heights, minIndex+1, r, tree))
	return res
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}*/

//  超出时间限制，时间复杂度 O(n^2)
func largestRectangleArea(heights []int) int {
	n := len(heights)
	res := 0
	for i := 0; i < n; i++ {
		// 找到高为 heights[i] 的面积最大的长方形，向前找到第一个高度小于 heights[i] 的位置，向后找到第一个高度大于 heights[i] 的位置
		w, h := 1, heights[i]
		j := i
		for j-1 >= 0 && heights[j-1] >= h {
			j--
			w++
		}
		j = i
		for j+1 < n && heights[j+1] >= h {
			j++
			w++
		}
		res = maxOf(res, w*h)
	}
	return res
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Test() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 4}))
}

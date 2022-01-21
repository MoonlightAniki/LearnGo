package solution0056

import (
	"fmt"
	"sort"
)

/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.



Example 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

Example 2:
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.


Constraints:
1 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 执行用时：656 ms, 在所有 Go 提交中击败了 5.29% 的用户
/*func merge(m [][]int) [][]int {
	intervals := Intervals{m}
	sort.Sort(&intervals)
	res := make([][]int, 0)
	solve(m, 0, &res)
	return res
}

func solve(intervals [][]int, index int, res *[][]int) {
	if index == len(intervals)-1 {
		*res = append(*res, intervals[index])
		return
	}
	solve(intervals, index+1, res)
	i := len(*res) - 1
	for ; i >= 0; i-- {
		if intervals[index][1] >= (*res)[i][0] {
			newRes := append([][]int{{intervals[index][0], maxOf((*res)[i][1], intervals[index][1])}}, (*res)[i+1:]...)
			*res = newRes
			break
		}
	}
	if i == -1 {
		newRes := append([][]int{intervals[index]}, *res...)
		*res = newRes
	}
}*/

// 执行用时：652 ms, 在所有 Go 提交中击败了 5.29% 的用户
/*func merge(m [][]int) [][]int {
	intervals := Intervals{m}
	sort.Sort(&intervals)
	return solve(m, 0)
}

func solve(intervals [][]int, index int) [][]int {
	if index == len(intervals)-1 {
		return [][]int{intervals[index]}
	}
	res := solve(intervals, index+1)
	i := len(res) - 1
	for ; i >= 0; i-- {
		if intervals[index][1] >= res[i][0] {
			res[i] = []int{intervals[index][0], maxOf(intervals[index][1], res[i][1])}
			return res[i:]
		}
	}
	return append([][]int{intervals[index]}, res...)
}*/

// 执行用时：20 ms, 在所有 Go 提交中击败了 31.97% 的用户
/*func merge(m [][]int) [][]int {
	intervals := Intervals{m}
	sort.Sort(intervals)
	var res [][]int
	res = append(res, m[0])
	for i := 1; i < len(m); i++ {
		if overlap(res[len(res)-1], m[i]) {
			res[len(res)-1] = mergeInterval(res[len(res)-1], m[i])
		} else {
			res = append(res, m[i])
		}
	}
	return res
}

func overlap(a []int, b []int) bool {
	return a[1] >= b[0]
}

func mergeInterval(a []int, b []int) []int {
	return []int{minOf(a[0], b[0]), maxOf(a[1], b[1])}
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Intervals struct {
	data [][]int
}

func (intervals Intervals) Len() int {
	return len(intervals.data)
}

func (intervals Intervals) Less(i, j int) bool {
	return intervals.data[i][0] < intervals.data[j][0]
}

func (intervals Intervals) Swap(i, j int) {
	intervals.data[i], intervals.data[j] = intervals.data[j], intervals.data[i]
}*/

// 执行用时：16 ms, 在所有 Go 提交中击败了 39.36% 的用户
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if temp[1] >= intervals[i][0] {
			temp[1] = maxOf(temp[1], intervals[i][1])
		} else {
			res = append(res, temp)
			temp = intervals[i]
		}
	}
	res = append(res, temp)
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
	fmt.Println(merge([][]int{
		{15, 18},
		{2, 6},
		{8, 10},
		{1, 3},
	}))
	fmt.Println(merge([][]int{
		{1, 4},
		{2, 3},
	}))
}

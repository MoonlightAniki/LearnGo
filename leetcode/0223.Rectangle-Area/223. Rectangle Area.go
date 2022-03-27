package solution0223

import "sort"

// 执行用时： 12 ms, 在所有 Go 提交中击败了 69.47% 的用户
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	areaA, areaB := (ax2-ax1)*(ay2-ay1), (bx2-bx1)*(by2-by1)
	if ax1 >= bx2 || ax2 <= bx1 || ay1 >= by2 || ay2 <= by1 {
		return areaA + areaB
	} else {
		x := []int{ax1, ax2, bx1, bx2}
		y := []int{ay1, ay2, by1, by2}
		sort.Ints(x)
		sort.Ints(y)
		overlapArea := (x[2] - x[1]) * (y[2] - y[1])
		return areaA + areaB - overlapArea
	}
}

package solution0315

import (
	"LearnGo/leetcode/utils"
	"testing"
)

func TestCountSmaller(t *testing.T) {
	testCases := []struct {
		nums   []int
		result []int
	}{
		{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
		{[]int{-1, -1}, []int{0, 0}},
	}
	for _, tc := range testCases {
		result := countSmaller(tc.nums)
		if !utils.SliceEquals(result, tc.result) {
			t.Errorf("Not Passed, nums:%v, expected result:%v, but got result:%v\n", tc.nums, tc.result, result)
		}
	}
}

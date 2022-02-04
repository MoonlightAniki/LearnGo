package solution0239

import (
	"LearnGo/leetcode/utils"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	testCases := []struct {
		nums   []int
		k      int
		result []int
	}{
		{
			nums:   []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:      3,
			result: []int{3, 3, 5, 5, 6, 7},
		},
		{
			nums:   []int{1},
			k:      1,
			result: []int{1},
		},
	}
	for _, tc := range testCases {
		result := maxSlidingWindow(tc.nums, tc.k)
		if !utils.SliceEquals(result, tc.result) {
			t.Errorf("Not Passed, nums:%v, k:%d, expected result:%v, but got result:%v\n", tc.nums, tc.k, tc.result, result)
		}
	}
}

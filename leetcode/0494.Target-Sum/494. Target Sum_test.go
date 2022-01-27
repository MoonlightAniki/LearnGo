package solution0494

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		result int
	}{
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{1}, 1, 1},
	}
	for _, tc := range testCases {
		result := findTargetSumWays(tc.nums, tc.target)
		if result != tc.result {
			t.Errorf("Not Passed, nums:%v, target:%v, expected result:%v, but got result:%v\n", tc.nums, tc.target, tc.result, result)
		}
	}
}

package solution0045

import "testing"

func TestJump(t *testing.T) {
	testCases := []struct {
		nums   []int
		result int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
	}
	for _, tc := range testCases {
		result := jump(tc.nums)
		if result != tc.result {
			t.Errorf("Not Passed, nums:%v, expected result:%d, but got result:%d\n", tc.nums, tc.result, result)
		}
	}
}

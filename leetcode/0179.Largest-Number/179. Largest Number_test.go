package solution0179

import "testing"

func TestLargestNumber(t *testing.T) {
	testCases := []struct {
		nums   []int
		result string
	}{
		{
			[]int{10, 2},
			"210",
		},
		{
			[]int{3, 30, 34, 5, 9},
			"9534330",
		},
	}
	for _, tc := range testCases {
		result := largestNumber(tc.nums)
		if result != tc.result {
			t.Errorf("Not Passed, nums:%v, expected result:%s, but got result:%s\n", tc.nums, tc.result, result)
		}
	}
}

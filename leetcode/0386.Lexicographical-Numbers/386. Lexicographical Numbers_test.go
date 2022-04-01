package solution0386

import (
	"LearnGo/leetcode/utils"
	"testing"
)

func TestLexicalOrder(t *testing.T) {
	testCases := []struct {
		n      int
		result []int
	}{
		{13, []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
		{2, []int{1, 2}},
	}
	for _, tc := range testCases {
		result := lexicalOrder(tc.n)
		if !utils.SliceEquals(result, tc.result) {
			t.Errorf("Not Passed, n:%d, expected result:%v, but got result:%v\n", tc.n, tc.result, result)
		}
	}
}

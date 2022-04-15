package solution0241

import (
	"LearnGo/leetcode/utils"
	"sort"
	"testing"
)

func TestDiffWaysToCompute(t *testing.T) {
	testCases := []struct {
		expression string
		result     []int
	}{
		{"2-1-1", []int{0, 2}},
		{"2*3-4*5", []int{-34, -14, -10, -10, 10}},
	}
	for _, tc := range testCases {
		result := diffWaysToCompute(tc.expression)
		sort.Ints(tc.result)
		sort.Ints(result)
		if !utils.SliceEquals(result, tc.result) {
			t.Errorf("Not Passed, expression:%s, expected result:%v, but got result:%v\n", tc.expression, tc.result, result)
		}
	}
}

package solution0001

import (
	"LearnGo/leetcode/utils"
	"testing"
)

// 函数名必须以 Test 开头
func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		res    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		if actual := twoSum(tt.nums, tt.target); !utils.SliceEquals(actual, tt.res) {
			t.Errorf("Not passed, nums:%v, target:%v, expected:%v, actual:%v\n", tt.nums, tt.target, tt.res, actual)
		}
	}
}

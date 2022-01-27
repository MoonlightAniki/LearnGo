package solution0322

import "testing"

func TestCoinChange(t *testing.T) {
	testCases := []struct {
		coins  []int
		amount int
		result int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		//{[]int{1}, 0, 0},
	}
	for _, tc := range testCases {
		result := coinChange(tc.coins, tc.amount)
		if result != tc.result {
			t.Errorf("Not Passed, coints:%v, amount:%v, got:%v, expected result:%v\n", tc.coins, tc.amount, result, tc.result)
		}
	}
}

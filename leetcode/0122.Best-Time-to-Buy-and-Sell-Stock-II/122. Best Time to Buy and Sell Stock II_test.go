package solution0122

import "testing"

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		prices []int
		result int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			result: 7,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			result: 4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			result: 0,
		},
	}
	for _, tc := range testCases {
		result := maxProfit(tc.prices)
		if result != tc.result {
			t.Errorf("Not Passed, prices:%v, expected result:%v, but got result:%v\n", tc.prices, tc.result, result)
		}
	}
}

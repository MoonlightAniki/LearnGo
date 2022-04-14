package solution0134

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	testCases := []struct {
		gas    []int
		cost   []int
		result int
	}{
		{
			[]int{2, 3, 4},
			[]int{3, 4, 3},
			-1,
		},
	}
	for _, tc := range testCases {
		result := canCompleteCircuit(tc.gas, tc.cost)
		if result != tc.result {
			t.Errorf("Not Passed, gas:%v, cost:%v, expected result:%d, but got result:%d\n", tc.gas, tc.cost, tc.result, result)
		}
	}
}

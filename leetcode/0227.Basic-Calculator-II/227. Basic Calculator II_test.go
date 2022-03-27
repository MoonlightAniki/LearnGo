package solution0227

import "testing"

func TestCalculate(t *testing.T) {
	testCases := []struct {
		s      string
		result int
	}{
		{"3+2*2", 7},
		{" 3/2 ", 1},
		{"3+5 / 2", 5},
	}
	for _, tc := range testCases {
		result := calculate(tc.s)
		if result != tc.result {
			t.Errorf("Not Passed, s:%s, expected result:%d, but got result:%d\n", tc.s, tc.result, result)
		}
	}
}

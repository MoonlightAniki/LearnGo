package solution0166

import "testing"

func TestFractionToDecimal(t *testing.T) {
	testCases := []struct {
		a      int
		b      int
		result string
	}{
		{1, 2, "0.5"},
		{2, 1, "2"},
		{4, 333, "0.(012)"},
		{-2147483648, 1, "-2147483648"},
	}
	for _, tc := range testCases {
		result := fractionToDecimal(tc.a, tc.b)
		if result != tc.result {
			t.Errorf("Not Passed, a:%d, b:%d, expected result:%s, but got result:%s\n", tc.a, tc.b, tc.result, result)
		}
	}
}

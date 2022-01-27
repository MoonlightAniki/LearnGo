package solution0279

import "testing"

func TestNumSquares(t *testing.T) {
	testCases := []struct {
		n     int
		count int
	}{
		{12, 3},
		{13, 2},
	}

	for _, tc := range testCases {
		if numSquares(tc.n) != tc.count {
			t.Errorf("Not passed, n = %d, count = %d, expected count = %d\n", tc.n, numSquares(tc.n), tc.count)
		}
	}
}

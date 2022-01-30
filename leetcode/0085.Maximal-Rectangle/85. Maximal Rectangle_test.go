package solution0085

import "testing"

func TestMaximalRectangle(t *testing.T) {
	testCases := []struct {
		matrix [][]byte
		result int
	}{
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			6,
		},
		{[][]byte{{'0'}}, 0},
		{[][]byte{{'1'}}, 1},
	}
	for _, tc := range testCases {
		result := maximalRectangle(tc.matrix)
		if result != tc.result {
			t.Errorf("Not Passed, matrix:%v, expected result:%v, but got result:%v\n", tc.matrix, tc.result, result)
		}
	}
}

package solution0130

import "testing"

func TestSolve(t *testing.T) {
	testCases := []struct {
		board  [][]byte
		result [][]byte
	}{
		{
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			result: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
	}
	for _, tc := range testCases {
		solve(tc.board)
		for i := range tc.board {
			for j := range tc.board[i] {
				if tc.board[i][j] != tc.result[i][j] {
					t.Errorf("Not Passed, expected result:%v, but got result:%v", tc.result, tc.board)
					return
				}
			}
		}
	}
}

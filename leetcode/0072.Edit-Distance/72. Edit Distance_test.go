package solution0072

import "testing"

func TestMinDistance(t *testing.T) {
	testCases := []struct {
		word1  string
		word2  string
		result int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}
	for _, tc := range testCases {
		result := minDistance(tc.word1, tc.word2)
		if result != tc.result {
			t.Errorf("Not Passed, word1:%v, word2:%v, expected result:%v, but got result:%v\n", tc.word1, tc.word2, tc.result, result)
		}
	}
}

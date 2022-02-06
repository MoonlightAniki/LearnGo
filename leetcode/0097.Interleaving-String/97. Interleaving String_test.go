package solution0097

import "testing"

func TestIsInterleave(t *testing.T) {
	testCases := []struct {
		s1, s2, s3 string
		result     bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
	}
	for _, tc := range testCases {
		result := isInterleave(tc.s1, tc.s2, tc.s3)
		if result != tc.result {
			t.Errorf("Not Passed, s1:%s, s2:%s, s3:%s, expected result:%v, but got result:%v\n", tc.s1, tc.s2, tc.s3, tc.result, result)
		}
	}
}

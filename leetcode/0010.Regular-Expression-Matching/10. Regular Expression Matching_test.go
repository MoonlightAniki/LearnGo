package solution0010

import "testing"

func TestIsMatch(t *testing.T) {
	testCases := []struct {
		s      string
		p      string
		result bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aaaa", "a*", true},
		{"abcdefg", ".*", true},
		{"aab", "c*a*b", true},
		{"a", "ab*", true},
		{"a", ".*...a", false},
	}
	for _, tc := range testCases {
		result := isMatch(tc.s, tc.p)
		if result != tc.result {
			t.Errorf("Not Passed, s:%v, p:%v, expected result:%v, but got result:%v\n", tc.s, tc.p, tc.result, result)
		}
	}
}

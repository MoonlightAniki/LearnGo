package solution0331

import "testing"

func TestIsValidSerialization(t *testing.T) {
	testCases := []struct {
		preorder string
		result   bool
	}{
		{"9,3,4,#,#,1,#,#,2,#,6,#,#", true},
		{"99,33,44,#,#,11,#,#,22,#,66,#,#", true},
		{"1,#", false},
		{"28,#,#", true},
		{"9,#,#,1", false},
		{"#", true},
	}
	for _, tc := range testCases {
		result := isValidSerialization(tc.preorder)
		if result != tc.result {
			t.Errorf("Not Passed, preorder:%s, expected result:%v, but got result:%v\n", tc.preorder, tc.result, result)
		}
	}
}

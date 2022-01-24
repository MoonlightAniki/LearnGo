package utils

import (
	"testing"
)

func TestSliceEquals(t *testing.T) {
	tests := []struct {
		a   []int
		b   []int
		res bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{1, 3, 2}, false},
		{[]int{}, nil, false},
		{[]int{}, []int{1}, false},
	}
	for _, tt := range tests {
		if actual := SliceEquals(tt.a, tt.b); actual != tt.res {
			t.Errorf("Not passed, SliceEquals(%v, %v) got %v, expected %v\n", tt.a, tt.b, actual, tt.res)
		}
	}
}

package popcount

import (
	"testing"
)

func TestPopcount(t *testing.T) {
	tests := []struct {
		input    uint64
		expected int
	}{
		{0, 0},
		{1, 1},
		{100, 3},
		{123456, 6},
	}
	for _, test := range tests {
		if actual := popcount(test.input); test.expected != actual {
			t.Errorf("input %b, expected %d, actual %d", test.input, test.expected, actual)
		}
	}
}

func TestSparsePopcount(t *testing.T) {
	tests := []struct {
		input    uint64
		expected int
	}{
		{0, 0},
		{1, 1},
		{100, 3},
		{123456, 6},
	}
	for _, test := range tests {
		if actual := sparsePopcount(test.input); test.expected != actual {
			t.Errorf("input %b, expected %d, actual %d", test.input, test.expected, actual)
		}
	}
}

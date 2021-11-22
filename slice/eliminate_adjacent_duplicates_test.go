package slice

import (
	"strings"
	"testing"
)

func TestEliminateAdjacentDuplicates(t *testing.T) {
	tests := []struct {
		in     []string
		expect []string
	}{
		{
			[]string{"a", "b", "c", "c", "c", "d", "d", "c", "e", "f", "f", "g"},
			[]string{"a", "b", "c", "d", "c", "e", "f", "g"},
		},
		{
			[]string{},
			[]string{},
		},
		{
			[]string{"a"},
			[]string{"a"},
		},
	}
	for _, test := range tests {
		inStr := "[" + strings.Join(test.in, " ") + "]"
		ret := EliminateAdjacentDuplicates(test.in)
		if strings.Join(test.expect, "") != strings.Join(ret, "") {
			t.Fatalf("input %s, expect %v, actual %v", inStr, test.expect, ret)
		}
	}
}

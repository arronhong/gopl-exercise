package slice

import (
	"strings"
	"testing"
)

func TestEliminateAdjacentDuplicates(t *testing.T) {
	tests := []struct {
		in     []string
		expect string
	}{
		{
			[]string{"a", "b", "c", "c", "c", "d", "d", "c", "e", "f", "f", "g"},
			"abcdcefg",
		},
		{
			[]string{},
			"",
		},
		{
			[]string{"a"},
			"a",
		},
	}
	for _, test := range tests {
		ret := EliminateAdjacentDuplicates(test.in)
		actual := strings.Join(ret, "")
		if test.expect != actual {
			t.Fatalf("expect %s, actual %s", test.expect, actual)
		}
	}
}

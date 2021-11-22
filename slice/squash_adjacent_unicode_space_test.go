package slice

import (
	"testing"
)

func TestSquashAdjacentUnicodeSpace(t *testing.T) {
	tests := []struct {
		in     []byte
		expect []byte
	}{
		{
			[]byte("hello  \t世 \n界"),
			[]byte("hello 世 界"),
		},
		{
			[]byte("hello  \t世 \n界  "),
			[]byte("hello 世 界 "),
		},
		{
			[]byte("hello  \t世 \n界\n\n"),
			[]byte("hello 世 界\n"),
		},
	}
	for _, test := range tests {
		inStr := string(test.in)
		ret := SquashAdjacentUnicodeSpace(test.in)
		if string(test.expect) != string(ret) {
			t.Fatalf("input %q, expect %q, actual %q", inStr, test.expect, ret)
		}
	}
}

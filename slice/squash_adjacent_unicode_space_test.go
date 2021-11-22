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
		ret := SquashAdjacentUnicodeSpace(test.in)
		actual := string(ret)
		expect := string(test.expect)
		if expect != actual {
			t.Fatalf("input %q, expect %q, actual %q", test.in, test.expect, ret)
		}
	}
}

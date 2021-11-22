package slice

import (
	"fmt"
	"testing"
)

func TestReverseUTF8EncodedBytes(t *testing.T) {
	tests := []struct {
		in     []byte
		expect []byte
	}{
		{
			[]byte("世界¢"),
			[]byte("¢界世"),
		},
		{
			[]byte{},
			[]byte{},
		},
		{
			[]byte("世¢"),
			[]byte("¢世"),
		},
		{
			[]byte("世h界¢"),
			[]byte("¢界h世"),
		},
		{
			[]byte("世"),
			[]byte("世"),
		},
	}
	for _, test := range tests {
		ret := ReverseUTF8EncodedBytes(test.in)
		fmt.Println(string(ret))
		if string(test.expect) != string(ret) {
			t.Fatalf("input %s, expect %s. actual %s", test.in, test.expect, ret)
		}
	}
}

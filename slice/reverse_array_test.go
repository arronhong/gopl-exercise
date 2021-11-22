package slice

import (
	"testing"
)

func TestReverseArray(t *testing.T) {
	s := [5]int{1, 2, 3, 4, 5}
	ReverseArray(&s)
	expect := [5]int{5, 4, 3, 2, 1}
	if expect != s {
		t.Fatalf("expect %v, actual %v", expect, s)
	}
}

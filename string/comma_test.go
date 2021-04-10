package goplstring

import (
	"testing"
)

func TestComma(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"12345", "12,345"},
		{"+12345", "12,345"},
		{"-12345", "-12,345"},
		{"12", "12"},
		{"+12", "12"},
		{"-12", "-12"},
		{"12345.678", "12,345.678"},
		{"+12345.678", "12,345.678"},
		{"-12345.678", "-12,345.678"},
		{"123", "123"},
		{"+123", "123"},
		{"-123", "-123"},
		{"123456", "123,456"},
		{"+123456", "123,456"},
		{"-123456", "-123,456"},
	}
	for _, test := range tests {
		if got := comma(test.s); got != test.want {
			t.Errorf("comma(%s) = %s, want %s", test.s, got, test.want)
		}
	}
}

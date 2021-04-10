package goplstring

import (
	"bytes"
	"strings"
)

// givin an input string representation of valid number, such as "12345", "123.45", "-12345.67"
// output string and insert commas every 3 place, as in "12,345", "123.45", "-123,45.67"
// no recursive and using bytes.Buffer instead of string concatenation
func comma(s string) string {
	// no input validation
	buf := bytes.Buffer{}
	if s[0] == '-' {
		buf.WriteByte('-')
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	f := ""
	lenOfInteger := strings.IndexByte(s, '.')
	if lenOfInteger < 0 {
		lenOfInteger = len(s)
	} else {
		f = s[lenOfInteger+1:]
	}
	for i, r := range s[:lenOfInteger] {
		if (lenOfInteger-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(r)
	}
	if len(f) > 0 {
		buf.WriteString("." + f)
	}
	return buf.String()
}

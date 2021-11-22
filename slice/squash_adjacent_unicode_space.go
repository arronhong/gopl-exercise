package slice

import (
	"unicode"
	"unicode/utf8"
)

func SquashAdjacentUnicodeSpace(bytes []byte) []byte {
	totSquashedSize := 0
	for i := 0; i < len(bytes)-totSquashedSize; {
		r1, size1 := utf8.DecodeRune(bytes[i:])
		if !unicode.IsSpace(r1) {
			i += size1
			continue
		}
		if i+size1 >= len(bytes)-totSquashedSize {
			break
		}
		r2, size2 := utf8.DecodeRune(bytes[i+size1:])
		if !unicode.IsSpace(r2) {
			i += size1 + size2
			continue
		}
		copy(bytes[i+size1:], bytes[i+size1+size2:])
		totSquashedSize += size2
	}
	return bytes[:len(bytes)-totSquashedSize]
}

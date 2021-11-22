package slice

import (
	"unicode/utf8"
)

func ReverseUTF8EncodedBytes(bytes []byte) []byte {
	for i, j := 0, len(bytes); ; {
		if i >= j {
			break
		}
		r1, size1 := utf8.DecodeRune(bytes[i:])
		if i+size1 >= j {
			break
		}
		r2, size2 := utf8.DecodeLastRune(bytes[i:j])
		middleSize := j - size2 - size1 - i
		if middleSize > 0 {
			copy(bytes[i+size2:], bytes[i+size1:i+size1+middleSize+1])
		}
		copy(bytes[i:], []byte(string(r2)))
		copy(bytes[i+size2+middleSize:], []byte(string(r1)))
		i += size2
		j -= size1
	}
	return bytes
}

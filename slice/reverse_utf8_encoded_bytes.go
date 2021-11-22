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

func Reverse(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

func ReverseUTF8EncodedBytes2(bytes []byte) []byte {
	// use the fact that: input []byte{x1, x2, x3, y1, y2, y3, y4, z1, z2}
	// if we reverse each runea then reverse whole bytes, we get the expection
	// after step1 like: []byte{x3, x2, x1, y4, y3, y2, y1, z2, z1}
	// after step2 like: []byte{z1, z2, y1, y2, y3, y4, x1, x2, x3}
	for i := 0; i < len(bytes); {
		_, size := utf8.DecodeRune(bytes[i:])
		Reverse(bytes[i : i+size])
		i += size
	}
	Reverse(bytes)
	return bytes
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	invalid := 0
	count := map[rune]int{}
	cate := map[string]int{}
	for {
		r, size, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && size == 1 {
			invalid++
			continue
		}
		count[r]++
		if unicode.IsControl(r) {
			cate["control"]++
		}
		if unicode.IsDigit(r) {
			cate["digit"]++
		}
		if unicode.IsGraphic(r) {
			cate["graphic"]++
		}
		if unicode.IsLetter(r) {
			cate["letter"]++
		}
		if unicode.IsMark(r) {
			cate["mark"]++
		}
		if unicode.IsNumber(r) {
			cate["number"]++
		}
		if unicode.IsPrint(r) {
			cate["print"]++
		}
		if unicode.IsPunct(r) {
			cate["punct"]++
		}
		if unicode.IsSpace(r) {
			cate["space"]++
		}
		if unicode.IsSymbol(r) {
			cate["symbol"]++
		}
		if unicode.IsUpper(r) {
			cate["upper"]++
		}
		if unicode.IsLower(r) {
			cate["lower"]++
		}
		if unicode.IsTitle(r) {
			cate["title"]++
		}
	}
	for c, n := range count {
		fmt.Printf("%q\t%d\n", c, n)
	}
	for ca, cnt := range cate {
		fmt.Printf("%s\t%d\n", ca, cnt)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "need file name on args\n")
		os.Exit(1)
	}
	counts := map[string]int{}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq open file error: %s\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	for s, cnt := range counts {
		fmt.Printf("%s\t%d\n", s, cnt)
	}
}

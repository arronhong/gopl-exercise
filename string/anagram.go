package goplstring

import (
	"reflect"
	"sort"
)

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	ss1 := []rune(s1)
	ss2 := []rune(s2)
	sort.Slice(ss1, func(i, j int) bool {
		return ss1[i] < ss1[j]
	})
	sort.Slice(ss2, func(i, j int) bool {
		return ss2[i] < ss2[j]
	})
	for i, r := range ss1 {
		if r != ss2[i] {
			return false
		}
	}
	return true
}

func anagram2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1 := map[rune]int{}
	m2 := map[rune]int{}
	for _, r := range s1 {
		m1[r]++
	}
	for _, r := range s2 {
		m2[r]++
	}
	return reflect.DeepEqual(m1, m2)
}

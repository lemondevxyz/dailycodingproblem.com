package main

import (
	"fmt"
	"strings"
)

func findPattern(w, s string) (ret []int) {
	start := 0
	add := 0

	for start != -1 {
		s = s[start:]

		index := strings.Index(s, w)
		if index == -1 {
			return
		}

		ret = append(ret, index+add)
		start = index + len(w)
		add += start
	}

	return
}

func reverse(s string) string {
	byt := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		byt = append(byt, s[i])
	}

	return string(byt)
}

func solve(w, s string) []int {
	return append(findPattern(w, s), findPattern(reverse(w), s)...)
}

func main() {
	fmt.Println("abxaba: ab", solve("ab", "abxaba"))
}

package main

import (
	"fmt"
	"strings"
)

func parse(str string) map[rune]int {
	m := map[rune]int{}

	split := strings.Split(str, "\n")
	for _, txt := range split {
		for _, v := range txt {
			m[v]++
		}
	}

	return m
}

func solve(word string, m map[rune]int) bool {
	for _, v := range word {
		_, ok := m[v]
		if !ok {
			return false
		} else {
			m[v]--

			if m[v] == 0 {
				delete(m, v)
			}
		}
	}

	return true
}

func main() {
	m := parse(`ABCE
SFCS
ADEE`)

	fmt.Println("ABCCED", solve("ABCCED", m))
	fmt.Println("SEE", solve("SEE", m))
	fmt.Println("ABCB", solve("ABCB", m))
}

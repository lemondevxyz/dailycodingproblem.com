package main

import "fmt"

func shift(a string) string {
	ret := []rune(a)

	return string(append(ret[1:], ret[0]))
}

func solve(a, b string) bool {
	if a == b {
		return true
	}

	for i := 0; i < len(b)+1; i++ {
		a = shift(a)
		if a == b {
			return true
		}
	}

	return false
}

func main() {
	a, b := "abcde", "cdeab"
	fmt.Println(a, b, solve(a, b))
	a, b = "abc", "acb"
	fmt.Println(a, b, solve(a, b))
}

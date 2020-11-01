// not gone lie, i was a bit lost at this one.
// for some reason the problem doesn't include a minimum length of substring. i.e substring could be as big as the string or smaller by one character.
// given the example i made substring length = k + 1
// time complexity is equal to O(n^2)
package main

import (
	"fmt"
	"os"
)

// int, int indiciates string index
func solve(s string, k int) string {

	if len(s) < k {
		return s
	}

	distinctchar := func(chars string) int {

		// index is character
		// value is occurance
		a := map[rune]bool{}

		for _, ruin := range []rune(chars) {
			a[ruin] = true
		}

		return len(a)

	}

	var start int
	var length = len(s)

	for i := 0; i < length; i++ {

		end := (i + k) + 1

		if end < length {

			dist := distinctchar(s[i:end])

			// stop at k distinct char
			if dist == k {
				start = i
				break
			}

		} else {
			break
		}
	}

	end := start + k + 1
	if end > len(s) {
		end = len(s)
	}

	return s[start:end]
}

var cases = map[string]string{
	solve("a", 2):        "a",
	solve("abcba", 2):    "bcb",
	solve("kacab", 2):    "aca",
	solve("toms1441", 5): "oms144",
	solve("ab", 2):       "ab",
}

func main() {

	fmt.Println("testing cases...")
	for k, v := range cases {
		if k != v {
			fmt.Println("bad code")
			// lamo as if this is going to help debugging hahaha
			// need to do data structure in the future
			fmt.Println(k, v)
			os.Exit(1)
		}
	}
	fmt.Println("success")

}

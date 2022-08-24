package main

import (
	"fmt"
	"os"
)

// const
var start = []rune{
	'[',
	'{',
	'(',
}

var end = []rune{
	']',
	'}',
	')',
}

func solve(str string) bool {

	next := []rune{}
	// range over characters
	for _, v := range str {
		index := -1
		// one of [, {, (
		for ks, vs := range start {
			if v == vs {
				index = ks
				break
			}
		}

		if index >= 0 {
			// prepend
			next = append([]rune{end[index]}, next...)
		} else {
			if len(next) > 0 {
				want := next[0]
				// one of ], }, )
				for ks, vs := range end {
					if v == vs {
						index = ks
						break
					}
				}

				// ignore if not end character
				if index >= 0 {
					// we closed it
					if want == end[index] {
						next = next[1:]
						continue
					} else {
						return false
					}
				}
			}
		}
	}

	// unclosed
	if len(next) > 0 {
		return false
	}

	return true
}

type Case struct {
	want bool
	have bool
}

var cases = []Case{
	{
		want: true,
		have: solve("([])[]({})"),
	},
	{
		want: false,
		have: solve("([)]"),
	},
	{
		want: false,
		have: solve("((()"),
	},
}

func main() {
	for k, v := range cases {
		if v.want != v.have {
			fmt.Printf("bad code. index: %d\nwant: %t have: %t", k, v.want, v.have)
			os.Exit(1)
		}
	}

	fmt.Println("success")
}

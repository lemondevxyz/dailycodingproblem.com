package main

import (
	"fmt"
	"os"
)

func solve(str string, regexp string) bool {
	// impossible
	if len(regexp) > len(str) {
		return false
	}

	// equality
	//fmt.Print("| ")
	match := 0

	regindex := 0
	for k := range str {
		vstr := str[k]
		peek := k + 1
		if peek >= len(str) {
			peek = -1
		}

		var vreg byte
		if regindex < len(regexp) {
			vreg = regexp[regindex]
			if vreg == vstr || vreg == '.' {
				match++
				regindex++
			} else if vreg == '*' {
				match = k + 1
				regindex++
			}
		}

		//fmt.Printf("%c - %c | ", vstr, vreg)
	}

	if match == len(str) {
		return true
	}

	return false
}

type Case struct {
	want bool
	have bool
}

var cases = []Case{
	{
		want: false,
		have: solve("raymond", "ra."),
	},
	{
		want: true,
		have: solve("ray", "ra."),
	},
	{
		want: false,
		have: solve("chats", ".*at"),
	},
	{
		want: true,
		have: solve("chat", ".*at"),
	},
}

func main() {
	for k, v := range cases {
		if v.want != v.have {
			fmt.Printf("bad code. index %d\n", k)
			os.Exit(1)
		}
	}
	fmt.Println("success")

}

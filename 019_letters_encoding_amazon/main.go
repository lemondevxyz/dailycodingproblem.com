package main

import (
	"fmt"
	"strings"
)

func encode(str string) string {
	if len(str) == 0 {
		return str
	}

	cur := str[0]
	occur := 0

	newstr := ""

	for i := 0; i < len(str); i++ {
		v := str[i]
		if v == cur {
			occur++
		} else {
			newstr = newstr + fmt.Sprintf("%d%c", occur, cur)
			occur = 1
			cur = str[i]
		}
	}
	// last piece
	newstr = newstr + fmt.Sprintf("%d%c", occur, cur)

	return newstr
}

func decode(str string) string {

	// skip two
	ret := ""
	for i := 1; i < len(str); i = i + 2 {
		occur := int(str[i-1] - 48)
		// letter
		if occur < 0 || occur > 9 {
			i = i - 1
			continue
		}

		char := str[i]
		ret = ret + strings.Repeat(string(char), occur)
	}

	return ret
}

func main() {
	ori := "AAAABBBCC"
	fmt.Println(ori, ":", encode(ori))

	ori = "6N9I1C0E"
	fmt.Println(ori, ":", decode(ori))
}

package main

import (
	"fmt"
	"strings"
)

func solve1(str string) (ret string) {
	split := strings.Split(str, " ")
	for i := len(split) - 1; i >= 0; i-- {
		ret += split[i] + " "
	}

	return ret
}

type MutableString []byte

func (m MutableString) String() string {
	return string([]byte(m))
}

func newMutable(str string) MutableString {
	return MutableString([]byte(str))
}

func reverse(str MutableString) MutableString {
	for i := 0; i < len(str)/2; i++ {
		j := len(str) - 1 - i

		str[i], str[j] = str[j], str[i]
	}

	return str
}

func solve2(str MutableString) string {
	for i := 0; i < len(str)/2; i++ {
		j := (len(str) - 1) - i

		str[i], str[j] = str[j], str[i]
	}

	lastIndex := -1
	for i := 0; i < len(str); i++ {
		if lastIndex == -1 {
			lastIndex = i
		}

		if str[i] == ' ' {
			reverse(str[lastIndex:i])
			lastIndex = -1
		}
	}
	reverse(str[lastIndex:])

	return str.String()
}

func main() {
	fmt.Println(solve1("hello world here"))
	fmt.Println(solve2(newMutable("hello world here")))
}

package main

import (
	"fmt"
	"strings"
)

const word_delimiter = " "

type StringSlice []string

func (s StringSlice) String() (str string) {
	str += "["
	for k, v := range s {
		// last
		str += v
		if k+1 < len(s) {
			str += ", "
		}
	}
	str += "]"

	return
}

// break it up
func solve(s string, k int) StringSlice {

	ret := StringSlice{}
	ret = append(ret, "")

	spl := strings.Split(s, word_delimiter)

	index := 0

	for i := 0; i < len(spl); i++ {
		word := spl[i]
		// how many chars in word
		wordsize := len(ret[index]) + len(word)
		// if it's bigger than k
		if wordsize > k {
			index++
		}

		if index >= len(ret) {
			ret = append(ret, "")
		}
		// there is a word
		if len(ret[index]) > 0 {
			ret[index] += " "
		}
		ret[index] += word
	}

	return ret
}

func main() {
	val := "the quick brown fox jumps over the lazy"
	k := 10

	fmt.Println(val, k, ":")
	fmt.Println(solve("the quick brown fox jumps over the lazy dog", k))
}

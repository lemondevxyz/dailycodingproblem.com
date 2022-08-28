package main

import (
	"fmt"
	"strings"
)

func slice(str string, set map[rune]struct{}) (ret []string, indexes map[int]struct{}) {
	indexes = map[int]struct{}{}
	last := 0

	for index, v := range str {
		_, exists := set[v]
		if exists {
			if len(str[last:index]) > 0 {
				ret = append(ret, str[last:index])
			}
			ret = append(ret, string(str[index]))
			indexes[len(ret)-1] = struct{}{}

			last = index + 1
		}
	}

	if len(str) > last {
		ret = append(ret, str[last:])
	}

	return
}

func reverse(str []string) {
	for i := 0; i < len(str)/2; i++ {
		j := len(str) - 1 - i

		str[i], str[j] = str[j], str[i]
	}
}

func solve(str string, set map[rune]struct{}) string {
	ret, indexes := slice(str, set)

	words := []string{}
	for index, val := range ret {
		if _, exists := indexes[index]; !exists {
			words = append(words, val)
		}
	}
	reverse(words)

	word := len(words) - 1
	for index := len(ret) - 1; index >= 0; index-- {
		if _, exists := indexes[index]; !exists {
			ret[index] = words[word]
			word--
		}
	}

	return strings.Join(ret, "")
}

func main() {
	fmt.Println(solve("hello/world:here", map[rune]struct{}{
		':': struct{}{},
		'/': struct{}{},
	}))
}

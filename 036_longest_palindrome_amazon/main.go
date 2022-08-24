package main

import "fmt"

func reverse(str string) string {
	ret := ""
	for _, v := range str {
		ret = string(v) + ret
	}

	return ret
}

func solve(str string) string {

	if str == reverse(str) {
		return str
	}

	for i := 0; i < len(str); i++ {
		// banana, 1
		// becomes anana
		substr := str[i:]
		if substr == reverse(substr) {
			return substr
		}

		j := (len(str) - 1) - i
		substr = str[i:j]
		// momo, 1
		// becomes mom
		if substr == reverse(substr) {
			return substr
		}

		// momo, 1
		// becomes mom
		substr = str[:j]
		if substr == reverse(substr) {
			return substr
		}
	}

	return ""
}

func main() {
	inputs := []string{"aabcdcb", "banana"}
	for _, v := range inputs {
		fmt.Printf("%s: %s\n", v, solve(v))
	}
}

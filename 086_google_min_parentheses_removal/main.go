package main

import "fmt"

func solve(str string) (result int) {
	lef := 0
	for _, v := range str {
		switch v {
		case '(':
			if lef < 0 {
				result++
			}
			lef++
		case ')':
			if lef == 0 {
				result++
			} else {
				lef--
			}
		}
	}

	if lef > 0 {
		result += lef
	}

	return result
}

func main() {
	fmt.Println("solve(\"()())()\"):", solve("()())()"))
	fmt.Println("solve(\")(\"):", solve(")("))
}

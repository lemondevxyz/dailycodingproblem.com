package main

import (
	"fmt"
)

func solve(str string) (num int) {
	lindex := map[int]bool{}
	rindex := map[int]bool{}

	for k, v := range str {
		if v == '(' {
			lindex[k] = true
			if rindex[k-1] {
				num++
			}
		} else if v == ')' {
			rindex[k] = true
			if lindex[k-1] {
				num--
			} else {
				num++
			}
		}
	}

	return num
}

func main() {
	str := "()())()"
	fmt.Println(str, solve(str))
	str = ")("
	fmt.Println(str, solve(str))
	str = ")()"
	fmt.Println(str, solve(str))
}

package main

import (
	"fmt"
	"strconv"
)

func conv(s string) uint8 {
	n, err := strconv.ParseUint(s, 2, 8)
	if err != nil {
		panic(err)
	}

	return uint8(n)
}

func swap(s string) string {
	by := []byte(s)
	for i := 0; i < len(by); i += 2 {
		j := i + 1
		if j < len(by) {
			by[i], by[j] = by[j], by[i]
		}
	}

	return string(by)
}

func solve(str string) string {
	return swap(str)
}

func main() {
	fmt.Println(solve("10101011"))
}

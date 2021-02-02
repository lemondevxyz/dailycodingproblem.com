package main

import "fmt"

// yo this problem is just amazing
func solve(x, y, b int32) int32 {
	return x*b + y*(1-b)
}

func main() {
	x, y, b := 32, 16, 1
	fmt.Printf("x: %d, y: %d, b: %d\n", x, y, b)
	fmt.Println(solve(32, 16, 1))
	x, y, b = 32, 16, 0
	fmt.Printf("x: %d, y: %d, b: %d\n", x, y, b)
	fmt.Println(solve(32, 16, 0))
}

package main

import "fmt"

// this works by testing the amount of minimums in an array
// if for example, there are 2 decreasing numbers - then there is no way to change on number cause it'll decrease anyway.
// while if there was one number, and the was the minimum. you could change numbers before it to make it increasing.
func solve(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}

	min := arr[0]
	count := 0
	for _, v := range arr[1:] {
		if v < min {
			min = v
			count++
		}

		if count >= 2 {
			return false
		}
	}

	if count == 1 {
		return true
	}

	return false
}

func main() {
	a := []int{10, 5, 7}
	fmt.Println(a, solve(a))
	a = []int{10, 5, 1}
	fmt.Println(a, solve(a))
}

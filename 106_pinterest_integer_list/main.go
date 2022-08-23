package main

import "fmt"

func solve(arr []int) bool {
	index := 0
	for {
		hop := arr[index]

		if hop == 0 {
			break
		}

		index += hop
		if index > len(arr) {
			index = len(arr) - 1
			break
		}
	}

	return index == len(arr)-1
}

func main() {
	arr := []int{2, 0, 1, 0}
	fmt.Println(arr, solve(arr))
	arr = []int{1, 1, 0, 1}
	fmt.Println(arr, solve(arr))
}

package main

import "fmt"

func sumSlice(slice []int) int {
	val := 0
	for _, v := range slice {
		val += v
	}

	return val
}

func wrap(arr []int, i, j int) []int {
	if j > len(arr) {
		start := 0
		end := j - len(arr)

		return append(arr[i:], arr[start:end]...)
	}

	return arr[i:j]
}

func solve(slice []int, want int) []int {
	for i := 1; i <= len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			arr := wrap(slice, j, j+i)

			if sumSlice(arr) == want {
				return arr
			}
		}
	}

	return nil
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(solve(arr, 9))
	fmt.Println(solve(arr, 6))
	fmt.Println(solve(arr, 5))
}

package main

import (
	"fmt"
	"math"
	"sort"
)

func solve(arr []int) float32 {

	//defer fmt.Println(arr)
	sort.Ints(arr)
	if len(arr) == 1 {
		return float32(arr[0])
	}

	size := float32(len(arr)) / 2
	round := math.Round(float64(size))

	even := false
	if size == float32(round) {
		even = true
	}

	if even {
		fir := int(size - 0.5)
		sec := int(size + 0.5)

		//fmt.Println(arr[fir], arr[sec])
		val := (float32(arr[fir] + arr[sec])) / 2

		return val
	}

	return float32(arr[int(size)])
}

func main() {
	input := []int{2, 1, 5, 7, 2, 0, 5}
	for i := 0; i < len(input); i++ {
		arr := input[:i+1]
		//fmt.Printf("%v ", arr)
		fmt.Print(solve(arr))
		fmt.Println()
	}
}

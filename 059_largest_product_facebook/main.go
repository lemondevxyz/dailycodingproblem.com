package main

import "fmt"

func solve(m int, arr []int) int {
	if m >= len(arr) {
		res := 0
		for _, v := range arr {
			if res == 0 {
				res = v
			} else {
				res = res * v
			}
		}

		return res
	}

	max := 0
	for i := 0; i < m; i++ {
		v := arr[i]
		if v == 0 {
			continue
		}

		if max == 0 {
			max = arr[i]
		} else {
			max = max * arr[i]
		}
	}

	for i := m; i < len(arr); i++ {

		div := arr[i-m]
		if div == 0 {
			continue
		}

		y := (max / div) * arr[i]

		if y > max {
			max = y
		}
	}

	return max
}

func main() {
	fmt.Println(solve(3, []int{-10, -10, 5, 2}))
}

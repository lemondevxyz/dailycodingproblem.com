package main

import "fmt"

// weak algorithm
/*
func algo(slice []int) {
	for i := len(slice) - 1; i >= 0; i-- {
		for j := len(slice) - 1; j >= 0; j-- {
			if i == j {
				continue
			} //
			fmt.Println(";", slice[i], slice[j])
			if slice[i] == (slice[j]-1) && j != 0 {
				for x := i; x >= 0; x-- { //
					if i < 0 || j < 0 {
						break
					}
					slice[i], slice[j] = slice[j], slice[i]
					i--
					j--
				}
				return
			} else if slice[i] < slice[j] && i > j {
				for x := 0; x < len(slice); x++ {
					if x == i {
						break
					}
					for y := x + 1; y < len(slice); y++ {
						slice[x], slice[y] = slice[y], slice[x]
						break
					}
				}
				return
			} else if slice[i] > slice[j] && i > j {
				slice[i], slice[j] = slice[j], slice[i]
				return
			}
		}
	}

	return
}
*/

func exp(x int) int {
	if x <= 1 {
		return 1
	} else {
		return x * exp(x-1)
	}
}

func add1(x, y int) int {
	if x+1 >= y {
		return 0
	}

	return x + 1
}

func swap(slice []int, i, j int) {
	if i >= len(slice) {
		return
	}

	for x := j + 1; x < len(slice); x++ {
		fmt.Println(i, j, x)
		if x+1 == len(slice) {
			slice[i], slice[j] = slice[j], slice[i]
			fmt.Println(slice)
			slice[i], slice[j] = slice[j], slice[i]
		}
		swap(slice, i, j+1)
	}

	swap(slice, i+1, j+1)
}

func solve(slice []int) {
	for i := 0; i < 1; i++ {
		iplus := i + 1
		if iplus >= len(slice) {
			iplus = 0
		}

		swap(slice, 0, iplus)
		//slice[0], slice[iplus] = slice[iplus], slice[0]
	}
}

func main() {
	// [1, 2, 3]: [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]
	slice := []int{1, 2, 3, 4}
	solve(slice)
}

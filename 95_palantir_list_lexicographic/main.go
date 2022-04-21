package main

import "fmt"

func subslice(slice []int, i int) []int {
	if i >= len(slice) { return append([]int{}, slice[:i]...) }
	return append(slice[:i], slice[i+1:]...)
}

// weak algorithm
func algo(slice []int) {
	for i := len(slice)-1; i >= 0; i-- {
		for j := len(slice)-1; j >= 0; j-- {
			if i == j { continue }//
				if slice[i] == (slice[j]-1) && j != 0 {
					//fmt.Println(";", slice[i], slice[j])
					for x := i; x >= 0; x-- {//
						if i < 0 || j < 0 { break }
						slice[i], slice[j] = slice[j], slice[i]
						i--
						j--
					}
					return
				} else if slice[i] < slice[j] && i > j {
					for x := 0; x < len(slice); x++ {
						if x == i { break }
						for y := x+1; y < len(slice); y++ {
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

func solve(slice []int) {
	fmt.Println("start", slice)
	for i := 0; i < (len(slice)*2)-1; i++ {
		algo(slice)
		fmt.Println(slice)
	}
	
	// reverse DaSlice
	for i := 0; i < len(slice); i++ {
		for j := i; j < len(slice); j++ {
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	fmt.Println("end", slice )
}

func main() {
	// [1, 2, 3]: [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]
	slice := []int{1, 2, 3}
	solve(slice)
}
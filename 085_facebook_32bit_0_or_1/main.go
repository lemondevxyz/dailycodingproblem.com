package main

import "fmt"

func solve(x, y, b int32) int32 {
	b += 1
	x = ((2 & b) / 2) * x
	y = ((1 & b) / 1) * y
	return x + y
}

func main() {
	for i := -100; i < 100; i++ {
		for j := -100; j < 100; j++ {
			for b := 0; b <= 1; b++ {
				val := solve(int32(i), int32(j), int32(b))
				if b == 1 && val != int32(i) {
					panic(fmt.Sprintf("something wrong: b = 1 but val isn't x. %d %d %d", i, j, b))
				} else if b == 0 && val != int32(j) {
					panic(fmt.Sprintf("something wrong: b = 1 but val isn't y. %d %d %d", i, j, b))

				}
			}
		}
	}

	x, y, b := int32(23), int32(-25), int32(1)
	fmt.Printf("solve(%d, %d, %d) = %d\n\nb = 0\n\n", x, y, b, solve(x, y, b))
	b = 0
	fmt.Printf("solve(%d, %d, %d) = %d\n", x, y, b, solve(x, y, b))
}

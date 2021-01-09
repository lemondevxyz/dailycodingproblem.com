package main

import "fmt"

// couldn't solve this one on my own :(
// rowcount, colcount, row current, col current
func traversal(rc, cc, ru, cu uint) uint {
	if ru == rc-1 && cu == cc-1 {
		return 1
	}

	count := uint(0)
	if ru < rc-1 {
		count += traversal(rc, cc, ru+1, cu)
	}
	if cu < cc-1 {
		count += traversal(rc, cc, ru, cu+1)
	}

	return count
}

func solve(size uint) uint {
	return traversal(size, size, 0, 0)
}

func main() {
	limit := 6
	for i := 1; i <= limit; i++ {
		fmt.Printf("%dx%d: %d\n", i, i, solve(uint(i)))
	}
}

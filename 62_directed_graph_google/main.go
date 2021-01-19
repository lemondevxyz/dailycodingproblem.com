package main

import "fmt"

type Place struct {
	src int
	dst int
}

func next(dst int, p []Place) int {
	for k, v := range p {
		if v.src == dst {
			return k
		}
	}

	return -1
}

func solve(p []Place) int {

	var dst int
	max := -1

	for _, v := range p {
		count := 0
		dst = v.dst

		if v.dst == v.src {
			continue
		}
		old := dst

		for {
			x := next(dst, p)
			if x == -1 {
				break
			} else {
				dst = p[dst].dst

				count++
			}
		}

		if dst != old {
			count++
		}

		if count > max {
			max = count
		}
	}

	return max
}

func main() {
	p := []Place{
		{0, 1},
		{0, 2},
		{2, 3},
		{3, 4},
	}

	//x := next(2, p)
	fmt.Println(solve(p))
	fmt.Println(solve([]Place{
		{0, 0},
	}))
}

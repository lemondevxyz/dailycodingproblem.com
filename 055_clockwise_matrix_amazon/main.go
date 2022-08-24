package main

import "fmt"

type Direction uint

const (
	Left Direction = iota + 1
	Right
	Up
	Down
)

func solve(m [][]int) {
	x, y := 0, 0

	minx := 0
	miny := 0

	maxy := len(m) - 1
	maxx := len(m[0]) - 1

	d := Right
	for k := 0; k < 20; k++ {
		fmt.Println(m[y][x])
		switch d {
		case Left:
			x--
			if x == minx {
				d = Up
				minx++
			}
		case Right:
			x++
			if x == maxx {
				d = Down
			}
		case Up:
			y--
			if y == miny {
				d = Right
			}
		case Down:
			y++
			if y == maxy {
				d = Left
				miny++
				maxy--
				maxx--
			}
		}
	}
}

func main() {
	matrix := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{6, 7, 8, 9, 10},
		[]int{11, 12, 13, 14, 15},
		[]int{16, 17, 18, 19, 20},
	}

	solve(matrix)
}

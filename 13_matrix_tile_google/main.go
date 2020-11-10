// i love this problem
// it's the perfect of difficulty where i push my problem-solving skills where it's not too eazy and not too hard.
// perfecto.
package main

import (
	"fmt"
	"os"
)

func solve(matrix [][]bool) (steps int) {

	steps = -1

	// bad
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	n := len(matrix)
	m := len(matrix[0])

	y := n - 1
	x := 0

	steps = 0
	// strategy
	// go through the down level first then try left step
	for {

		peek := y - 1
		if peek < 0 {
			peek = 0
		}

		// original y
		// oy := y
		// original x
		ox := x

		if !matrix[peek][x] && y > 0 {
			y--
			// couldn't find it :(
			if y < 0 {
				return -1
			}
		} else {
			if y > 0 {
				x++
				// couldn't find it :(
				if x >= m {
					return -1
				}
			} else {
				x--
				// x = 0
				// y = 0
				// success
				if x < 0 {
					x = 0
					break
				}
			}
		}

		v := matrix[y][x]
		if !v {
			steps++
		} else {
			// mark that as not avaliable
			// dead end
			matrix[y][x] = true
			matrix[y][ox] = true

			// go down a level
			y = y + 1
			x = ox
		}
	}

	return
}

type Case struct {
	want int
	have int
}

var cases = []Case{
	{
		have: solve([][]bool{
			0: {false, false, false, false},
			1: {true, true, false, true},
			2: {false, false, false, false},
			3: {false, false, false, false},
		}),
		want: 7,
	},
	{
		have: solve([][]bool{
			0: {false, false, false, false},
			1: {false, false, false, false},
			2: {false, false, false, false}}),
		want: 2,
	},
	{
		have: solve([][]bool{
			0: {false, false, false, false}}),
		want: 0,
	},
	{
		have: solve([][]bool{
			0: {false, false, false, false},
			1: {true, true, true, false},
			2: {false, true, true, false},
			3: {false, false, false, false},
		}),
		want: 10,
	},
	{
		have: solve([][]bool{
			0: {false, false, false, true},
			1: {true, true, true, false},
			2: {false, true, true, false},
			3: {false, false, false, false},
		}),
		want: -1,
	},
}

func main() {

	for k, v := range cases {
		if v.have != v.want {
			fmt.Println("bad code. index is", k)
			fmt.Println(v.want, v.have)
			os.Exit(1)
		}
	}

	fmt.Println("success")
}

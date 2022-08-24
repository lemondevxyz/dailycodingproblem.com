package main

import (
	"fmt"
	"strings"
)

type Matrix [][]bool

func (m Matrix) String() (str string) {
	ylen := len(m)
	xlen := 0
	if ylen > 0 {
		xlen = len(m[0])
	}

	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			b := m[y][x]

			if b {
				str += "1"
			} else {
				str += "0"
			}

			if x+1 < xlen {
				str += " "
			}
		}

		if y+1 < ylen {
			str += "\n"
		}
	}

	return
}

type Point struct {
	Y int
	X int
}

func (p Point) String() string {
	return fmt.Sprintf("%d:%d", p.Y, p.X)
}

func (p Point) Square() []Point {
	return []Point{
		{p.Y - 1, p.X - 1},
		{p.Y - 1, p.X},
		{p.Y - 1, p.X + 1},
		{p.Y, p.X - 1},
		{p.Y, p.X + 1},
		{p.Y + 1, p.X - 1},
		{p.Y + 1, p.X},
		{p.Y + 1, p.X + 1},
	}
}

func (p Point) CheckForSquare(m Matrix, mt map[string]bool) (ret []Point) {
	if p.Y < 0 || p.X < 0 || p.Y >= len(m) || p.X >= len(m[0]) {
		return nil
	}

	_, ok := mt[p.String()]
	if m[p.Y][p.X] && !ok {
		mt[p.String()] = true
		ret = append(ret, p)

		ps := p.Square()
		for _, v := range ps {
			ret = append(ret, v.CheckForSquare(m, mt)...)
		}
	}

	return ret
}

func parseMatrix(str string) Matrix {
	yarr := strings.Split(str, "\n")

	m := Matrix{}
	for y := 0; y < len(yarr); y++ {
		xarr := strings.Split(yarr[y], " ")
		m = append(m, []bool{})
		for x := 0; x < len(xarr); x++ {
			ch := xarr[x]
			b := false
			if ch == "1" {
				b = true
			}

			m[y] = append(m[y], b)
		}
	}

	return m
}

func solve(str string) (islands int) {
	m := parseMatrix(str)

	for y := range m {
		for x := range m[y] {
			//v := m[y][x]

			ctx := map[string]bool{}

			p := Point{y, x}
			ps := p.CheckForSquare(m, ctx)

			for _, v := range ps {
				m[v.Y][v.X] = false
			}

			if len(ps) > 0 {
				islands++
			}
		}
	}

	return
}

// strategy
// convert to slice so its easier to deal with
// look for any booleans and surrounding ones
// add em to a map so that crossing them again does not happen
// then remove the acutal points from the matrix
func main() {
	m := solve(`1 0 0 0 0
0 0 1 1 0
0 1 1 0 0
0 0 0 0 0
1 1 0 0 1
1 1 0 0 1`)

	fmt.Println(m)
}

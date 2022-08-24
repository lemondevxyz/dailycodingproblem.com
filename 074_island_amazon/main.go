package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Matrix [][]bool

const templ = "%02d%02d"

func abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func (p Point) Square() []Point {
	return []Point{
		{p.X + 1, p.Y + 1},
		{p.X, p.Y + 1},
		{p.X - 1, p.Y + 1},
		{p.X + 1, p.Y},
		// {p.X, p.Y}, our original point
		{p.X - 1, p.Y},
		{p.X + 1, p.Y - 1},
		{p.X, p.Y - 1},
		{p.X - 1, p.Y - 1},
	}
}

func NewMatrix(x, y int) Matrix {
	m := make(Matrix, y)
	for k := range m {
		m[k] = make([]bool, x)
	}

	return m
}

func (m Matrix) String() (str string) {
	for k, s := range m {
		if k != 0 {
			str += "\n"
		}

		for _, v := range s {
			num := 0
			if v {
				num = 1
			}
			str += fmt.Sprintf("%d ", num)
		}
	}

	return
}

func (m Matrix) Set(z []Point) {
	for _, s := range z {
		m[s.X][s.Y] = true
	}
}

func (m Matrix) GoAround(point Point, exist map[string]bool) {

	exist[fmt.Sprintf(templ, point.X, point.Y)] = true

	for _, v := range point.Square() {
		// not of bounds
		if v.X >= 0 && v.Y >= 0 && v.X < len(m) && v.Y < len(m[0]) {

			if m[v.X][v.Y] {
				if !exist[fmt.Sprintf(templ, v.X, v.Y)] {
					m.GoAround(v, exist)
				}
			}
		}
	}
}

func (m Matrix) Island() (num int) {

	exist := map[string]bool{}

	for x, v := range m {
		for y, s := range v {
			if s && !exist[fmt.Sprintf(templ, x, y)] {
				m.GoAround(Point{x, y}, exist)
				num++
			}
		}
	}

	return
}

func main() {
	m := NewMatrix(5, 6)
	m.Set([]Point{
		{0, 0},
		{1, 2},
		{1, 3},
		{2, 1},
		{2, 2},
		{4, 0},
		{4, 1},
		{4, 4},
		{5, 0},
		{5, 1},
		{5, 4},
	})

	fmt.Println("island")
	fmt.Println(m)
	fmt.Println()
	fmt.Println(m.Island())
	//fmt.Println(Point{4, 0}.Square())
}

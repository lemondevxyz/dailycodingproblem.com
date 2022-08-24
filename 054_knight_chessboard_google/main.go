package main

import "fmt"

type chess [][]int

type cord struct {
	x int
	y int
}

func (c chess) String() string {
	str := ""
	for k, v := range c {
		for _, s := range v {
			str += fmt.Sprintf("%02d ", s)
		}

		if k+1 < len(c) {
			str += "\n"
		}
	}

	return str
}

func newBoard(n uint) chess {
	c := make(chess, n)
	for i := uint(0); i < n; i++ {
		c[i] = make([]int, n)
	}

	return c
}

func (c chess) Where() cord {
	for i, s := range c {
		for k, v := range s {
			if v == 1 {
				return cord{x: i, y: k}
			}
		}
	}

	return cord{x: -1, y: -1}
}

func (c chess) Largest() int {
	x := 1

	for _, s := range c {
		for _, v := range s {
			if v > x {
				x = v
			}
		}
	}

	return x
}

func (c chess) Potential() []cord {

	wh := c.Where()
	x, y := wh.x, wh.y

	if x == -1 || y == -1 {
		return nil
	}

	cords := []cord{
		cord{x: x - 2, y: y - 1},
		cord{x: x + 2, y: y - 1},
		cord{x: x - 2, y: y + 1},
		cord{x: x + 2, y: y + 1},
		cord{x: x - 1, y: y - 2},
		cord{x: x - 1, y: y + 2},
		cord{x: x + 1, y: y - 2},
		cord{x: x + 1, y: y + 2},
	}

	for i := len(cords) - 1; i >= 0; i-- {
		v := cords[i]
		if v.y < 0 || v.x < 0 || v.x >= len(c) || v.y >= len(c) {
			cords = append(cords[:i], cords[i+1:]...)
		}
	}

	return cords
}

// recursive function
func (c chess) Run() {

	oc := c.Where()
	if oc.x == -1 || oc.y == -1 {
		return
	}
	ocs := c.Potential()

	i := 0

	for _, v := range ocs[:1] {
		tempc := chess{}
		for _, s := range c {
			sl := []int{}
			for _, v := range s {
				sl = append(sl, v)
			}

			tempc = append(tempc, sl)
		}

		tempc[oc.x][oc.y] = tempc.Largest() + 1
		tempc[v.x][v.y] = 1

		tempc.Run()
		fmt.Println(tempc, "\n")

		if i == 25 {
			break
		}
	}

	return
}

func solve(n uint) uint {

	b := newBoard(5)
	b[2][2] = 1
	b.Run()

	return 0
}

func main() {
	solve(5)
}

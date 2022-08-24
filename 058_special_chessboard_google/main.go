package main

import "fmt"

type Place struct {
	x int
	y int
}

type Direction uint8

const (
	UpLeft Direction = iota + 1
	UpRight
	DownLeft
	DownRight
)

func (d Direction) Valid() bool {
	return d >= UpLeft && d <= DownRight
}

// Time O(n^2), Space O(n)
func solve(m uint, p []Place) (res int) {
	exist := map[int]bool{}
	for _, c := range p {
		x, y := c.x, c.y
		if x >= int(m) || y >= int(m) {
			continue
		}

		dir := UpLeft
		for {
			if dir == UpLeft {
				x++
				y--
			} else if dir == UpRight {
				x++
				y++
			} else if dir == DownLeft {
				x--
				y--
			} else if dir == DownRight {
				x--
				y++
			}

			if x > int(m) || y > int(m) || y < 0 || x < 0 {
				dir++
			} else {
				for k, v := range p {
					if v.x != c.x && v.y != c.y {
						if v.x == x && v.y == y {
							if _, ok := exist[k]; !ok {
								exist[k] = true
								res++
								break
							}
						}
					}
				}
			}

			if !dir.Valid() {
				break
			}
		}
	}

	return
}

func main() {
	p := []Place{
		{0, 0},
		{1, 2},
		{2, 2},
		{4, 0},
	}
	fmt.Println(p, ":", solve(5, p))
}

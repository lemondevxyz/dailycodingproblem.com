package main

import "fmt"

type Tower [3][]int

func (t Tower) String(max int) string {
	str := ""
	for _, v := range t {
		if len(v) > max {
			max = len(v)
		}
	}

	for i := 0; i < max; i++ {
		for _, v := range t {
			index := max - 1 - i
			if len(v) > index {
				str += fmt.Sprint(v[index], " ")
			} else {
				str += fmt.Sprint(".", " ")
			}
		}

		str += fmt.Sprint("\n")
	}

	return str[:len(str)-1]
}

func (t *Tower) Move(src, dst int) bool {
	srcArr := t[src]
	dstArr := t[dst]

	if len(srcArr) == 0 {
		return false
	}
	srcVal := srcArr[len(srcArr)-1]

	var dstVal *int
	if len(dstArr) > 0 {
		dstVal = &dstArr[len(dstArr)-1]
	}

	if dstVal != nil && *dstVal < srcVal {
		//panic(fmt.Sprintf("%d should be smaller than %d", *dstVal, srcVal))
		return false
	}

	t[dst] = append(dstArr, srcArr[len(srcArr)-1])
	t[src] = t[src][:len(srcArr)-1]

	return true
}

func pow(x, y int) int {
	res := 1
	for i := 0; i < y; i++ {
		res = res * x
	}

	return res
}

type move struct {
	src int
	dst int
}

func (m move) String() string {
	return fmt.Sprintf("Move %d to %d", m.src+1, m.dst+1)
}

// shoutout to wikipedia
func solve(t *Tower) []move {
	moves := []move{}

	disks := len(t[0])
	isOdd := (disks % 2) == 1
	step := 0

	for {
		if len(t[2]) == disks {
			break
		}

		switch step {
		case 0:
			if isOdd {
				if !t.Move(0, 2) {
					t.Move(2, 0)
					moves = append(moves, move{2, 0})
				} else {
					moves = append(moves, move{0, 2})
				}
			} else {
				if !t.Move(0, 1) {
					t.Move(1, 0)
					moves = append(moves, move{1, 0})
				} else {
					moves = append(moves, move{0, 1})
				}
			}
			step++
		case 1:
			if isOdd {
				if !t.Move(0, 1) {
					moves = append(moves, move{1, 0})
					t.Move(1, 0)
				} else {
					moves = append(moves, move{0, 1})
				}
			} else {
				if !t.Move(0, 2) {
					t.Move(2, 0)
					moves = append(moves, move{2, 0})
				} else {
					moves = append(moves, move{0, 2})
				}
			}
			step++

		case 2:
			if !t.Move(1, 2) {
				moves = append(moves, move{2, 1})
				t.Move(2, 1)
			} else {
				moves = append(moves, move{1, 2})
			} // B C
			step = 0
			break
		}
	}

	return moves
}

func main() {
	t := &Tower{{3, 2, 1}}

	moves := solve(t)
	for _, move := range moves {
		fmt.Println(move)
	}
}

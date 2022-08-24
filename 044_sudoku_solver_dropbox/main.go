package main

import (
	"fmt"
	"math"
)

type Board [][]int

func (b Board) GetMod(i int) int {
	size := len(b)
	a := b.BoxAmount()

	for j := 0; j < size; j += a {
		if i >= j && i <= (j+(a-1)) {
			return j
		}
	}

	return -1
}

// amount of boxes in one row
func (b Board) BoxAmount() int {
	return int(math.Sqrt(float64(len(b))))
}

// O(n^2) time complexity, O(n) space
func (b Board) Valid() bool {

	if len(b) == 0 {
		return false
	} else if len(b) > 0 && len(b) != len(b[0]) {
		return false
	}

	size := len(b)

	// row column box check
	x := 0
	for i := 0; i < size; i++ {
		row := map[int]bool{}
		col := map[int]bool{}

		for j := 0; j < size; j++ {
			row[b[i][j]] = true
			col[b[j][i]] = true
		}

		box := b.GetBox(x)
		occ := map[int]bool{} // short for occur
		for _, v := range box {
			if _, ok := occ[v]; !ok {
				occ[v] = true
			} else {
				return false
			}
		}

		if len(row) != size || len(col) != size {
			return false
		}
	}

	return true
}

func (b Board) String() string {
	str := ""

	for k, v := range b {
		islast := (k + 1) > len(b)
		str += fmt.Sprintf("%v", v)
		if !islast {
			str += "\n"
		}
	}

	return str
}

// O(n) time complexity, constant space
// where i is the box pos
func (b Board) GetBox(i int) (slice []int) {
	a := b.BoxAmount()

	x := i % a
	x = x * a

	i = b.GetMod(i)
	slice = make([]int, len(b))

	for j := i; j < i+a; j++ {
		copy(slice, b[i][x:x+a])
	}

	return
}

func (b Board) ReplaceBox(i int, box []int) {
	size := len(b)

	a := b.BoxAmount()

	x := i % a
	x = x * a
	//fmt.Println(y, i)
	i = b.GetMod(i)

	//x = 0
	for j := 0; j < size; j += a {
		start := j
		end := start + a
		copy(b[i][x:(x+a)], box[start:end])
		i++
	}
}

func NewBoard(size int) Board {
	b := make(Board, size)
	for i := 0; i < size; i++ {
		b[i] = make([]int, size)
	}

	return b
}

// solve by creating first box, then manipulate order and add to board
// O(n^2)
func solve(b Board) {

	// get the first box
	box := b.GetBox(0)
	// which numbers are duplicates
	exist := map[int]bool{}
	size := len(b)

	// which indexes to replace
	zeros := []int{}

	// get where are the zeros
	for k, v := range box {
		_, ok := exist[v]
		if !ok {
			if v != 0 {
				exist[v] = true
			} else {
				zeros = append(zeros, k)
			}
		} else {
			// duplicate
			return
		}
	}

	// fill zeros in the box with number
	for i := 0; i < len(zeros); i++ {
		v := zeros[i]
		for j := 1; j <= size; j++ {
			_, ok := exist[j]
			if !ok {
				box[v] = j
				exist[j] = true
				break
			}
		}
	}

	add := int(math.Sqrt(float64(size)))
	max := add - 1

	// manipulate row
	manirow := func() []int {
		return append(box[add:], box[:add]...)
	}
	// manipulate column
	manicol := func() []int {
		newbox := []int{}
		for i := 0; i < len(box); i += add {
			start := i
			end := i + max + 1

			// doing this:
			// slice = box[start:end]
			// wouldn't produce unwanted results
			slice := make([]int, add)
			copy(slice, box[start:end])
			slice = append(slice[1:], slice[0])

			newbox = append(newbox, slice...)
		}

		return newbox
	}

	b.ReplaceBox(0, box)

	// start manipulating box order
	for i := 1; i < size; i++ {
		//x := i + 1
		if (i % add) == 0 {
			box = manicol()
		} else {
			box = manirow()
		}

		b.ReplaceBox(i, box)
	}

}

// O(n^2)
func main() {
	// creates empty board with zeros
	b := NewBoard(9)
	b[0][0] = 2
	b[0][1] = 5
	b[0][2] = 9
	b[1][2] = 3
	b[2][2] = 4
	b[2][1] = 1
	//b[0][5] = 6
	b[3][0] = 9

	fmt.Println(b)
	solve(b)
	fmt.Println(b)

}

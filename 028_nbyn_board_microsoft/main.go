package main

import "fmt"

type board [][]bool

func (b board) String() (str string) {
	for _, v := range b {
		for _, c := range v {
			char := "o"
			if c {
				char = "x"
			}

			str += fmt.Sprintf("%s ", char)
		}
		str += "\n"
	}

	return
}

// true if it's invalid
// false if it's valid
func (b board) safe() bool {
	// i = j
	// bottom-left
	sum := map[int]bool{}
	// top-right
	sub := map[int]bool{}
	// column
	col := map[int]bool{}
	for i, v := range b {
		trues := 0
		for j, c := range v {
			if c {
				if _, ok := col[j]; ok {
					return false
				}

				if _, ok := sum[i+j]; ok {
					return false
				}

				if _, ok := sub[i-j]; ok {
					return false
				}

				col[j] = true
				sum[i+j] = true
				sub[i-j] = true

				trues++

				if trues > 1 {
					return false
				}
			}
		}
	}

	return true
}

func solve(n int, start int) board {

	b := board{}
	for i := 0; i < n; i++ {
		b = append(b, []bool{})
		k := len(b) - 1
		for j := 0; j < n; j++ {

			b[k] = append(b[k], true)

			if !b.safe() || (i == 0 && j != start) {
				b[k][j] = false
			}
		}
	}

	return b
}

// queens = N <= 4 ? n-2 : n
func main() {
	fmt.Println(solve(3, 2))
}

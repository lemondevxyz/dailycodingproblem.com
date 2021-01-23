package main

import (
	"fmt"
	"strings"
)

var (
	str = []byte("abcdefghijklmnopqrstuvwxyz")
)

func get_index(b byte) int {
	for k, v := range str {
		if v == b {
			return k
		}
	}

	return -1
}

func str_to_table(str string) table {
	spl := strings.Split(str, "\n")
	t := table{}
	for x, v := range spl {
		t = append(t, make([]byte, len(v)))
		for y, c := range []byte(v) {
			t[x][y] = c
		}
	}

	return t
}

type table [][]byte

func (t table) String() (str string) {
	for k, s := range t {
		if k != 0 {
			str += "\n"
		}

		for _, v := range s {
			str += string(v)
		}
	}

	return str
}

func solve(t table) (count int) {

	max := map[int]int{}
	skip := map[int]bool{}

	for x, s := range t {
		for y, v := range s {
			if skip[y] {
				continue
			}

			i := get_index(v)
			if x == 0 && y == 0 {
				max[y] = i
				continue
			}

			o := max[y]
			if o <= i {
				max[y] = i
			} else if o > i {
				count++
				skip[y] = true
			}
		}
	}

	return count
}

func main() {
	str := `cba
daf
ghi`
	fmt.Println(str, solve(str_to_table(str)))
	fmt.Println()

	str = "abcdefg"
	fmt.Println(str, solve(str_to_table("abcdefg")))
	fmt.Println()

	str = `zyx
wvu
tsr`
	fmt.Println(str, solve(str_to_table(str)))
}

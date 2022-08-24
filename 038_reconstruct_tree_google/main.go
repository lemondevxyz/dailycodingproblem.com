package main

import (
	"fmt"
	"strings"
)

func get_pos(arr []string, want string) int {
	for k, v := range arr {
		if v == want {
			return k
		}
	}

	return -1
}

// since how the string is encoded is not clear, this function
// will only work with 7-length slices
// no-cap i definitely was tripping on this problem.
func solve(torder, forder []string) (str string) {

	basepos := get_pos(forder, torder[0])
	levels := [][]string{}

	for tk, tv := range torder {
		fk := get_pos(forder, tv)
		if fk == -1 {
			return
		}

		diff := basepos - fk
		if fk > basepos {
			diff = fk - basepos
		}

		level := 2
		if diff == 0 {
			// a
			level = 0
		} else if diff == 2 {
			// b, c
			level = 1
		}

		if len(levels) == level {
			levels = append(levels, []string{
				tv,
			})
		}

		// tabs > spaces
		tabs := strings.Repeat(" ", level)
		str += tabs + fmt.Sprintf("- %s", tv)
		if tk+1 != len(torder) {
			str += "\n"
		}
	}

	return
}

// a
// - b
//   - d
//   - e
// - c
//   - f
//   - g

// d
// - b
// e
//   - a
// f
// - c
// g
func main() {
	torder := []string{"a", "b", "d", "e", "c", "f", "g"}
	forder := []string{"d", "b", "e", "a", "f", "c", "g"}

	// it seems my programming skills don't include drawing a tree
	fmt.Println(solve(torder, forder))
}

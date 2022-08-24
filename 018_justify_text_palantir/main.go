package main

import (
	"fmt"
	"sort"
	"strings"
)

// i.e sort by increasing length
type lenStrings []string

func (l lenStrings) Len() int {
	return len(l)
}

func (l lenStrings) Less(i, j int) bool {
	return len(l[i]) > len(l[j])
}

func (l lenStrings) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// O(n^2), some optimization is done so that it breaks asap.
func get_optimal_length(length int) int {

	results := []int{}

	for i := 2; i < length; i++ {
		for k := 2; k < length; k++ {
			mul := i * k
			if mul > length {
				// all hope is lost
				break
			} else if mul == length {
				// make our latest multiplication the first in arr
				results = append([]int{
					i,
				}, results...)
			}
		}
	}

	if len(results) > 0 {
		return results[0]
	}

	return 1
}

// O(n^2 + 2n + n*m)
// get_optimal_length() is O(n^2)
// createLines() is O(n) and is called twice.
func solve(arg []string) []string {
	// prevent slice modification
	words := []string{}
	words = append(words, arg...)

	// fit as many words into as little lines as possible
	m := get_optimal_length(len(words))

	createLines := func() (lines []string) {
		for k := range words {

			rel := k % m // relative index
			if rel == 0 {
				end := k + m
				lines = append(lines, strings.Join(words[k:end], " "))
			}

		}

		return lines
	}

	lines := lenStrings(createLines())
	if len(lines) == 0 {
		return lines
	}

	// sorting lines makes sure that lines[0] is the biggest one.
	sort.Sort(lines)
	want := len(lines[0])

	padright := false
	if m == 1 {
		// meaning 1 per word per line
		padright = true
	}

	if !padright {
		for k, v := range lines {

			have := len(v)
			if have == want {
				// no need to change
				continue
			}

			ind := k * m

			// want, and have split
			// to compare words
			j := 0

			for i := ind; i < ind+m; i++ {
				// having := does not change the above variables
				want := len(words[j])
				have := len(words[i])

				if want > have {
					diff := want - have
					add := strings.Repeat(" ", diff)

					words[i] = add + words[i]
				}
				j++

			}
		}
	} else {
		for i := range words {
			v := words[i]
			have := len(v)

			diff := want - have

			words[i] = words[i] + strings.Repeat(" ", diff)
		}
	}

	return createLines()
}

func prettyprint(lines []string) {
	for _, v := range lines {
		fmt.Print(v)
		if strings.HasSuffix(v, " ") {
			fmt.Print(len(v))
		}

		fmt.Println()
	}
}

func main() {
	val := []string{
		"the", "quick", "brown",
		"fox", "jumps", "over",
		"the", "lazy", "dog",
	}

	prettyprint(solve(val))

	fmt.Println()

	val = append(val, "ayay", "ayay")
	prettyprint(solve(val))
}

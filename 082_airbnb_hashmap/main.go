package main

import (
	"fmt"
	"math"
)

func solve(m map[string][]string) []string {
	base := ""

	for c, v := range m {
		if len(v) == 0 {
			base = c
			break
		}
	}

	sli := []string{}

	for i := 0; i < len(m); i++ {
		oldbase := base
		min := math.MaxInt64

		sli = append(sli, base)

		for k, v := range m {
			if k == base {
				continue
			}

			for _, s := range v {
				if s == oldbase {
					if len(v) < min {
						base = k
						min = len(v)
					}
				}
			}
		}
	}

	return sli
}

func prettyprint(x map[string][]string) {
	for k, v := range x {
		fmt.Println(k, ":", v)
	}
}

func main() {
	m := map[string][]string{
		"zxc": {"vbn", "ghl", "asd"},
		"vbn": {"ghl", "asd"},
		"ghl": {"asd"},
		"asd": nil,
	}

	prettyprint(m)
	fmt.Println(solve(m))

	m = map[string][]string{
		"CSC300": {"CSC200", "CSC100"},
		"CSC200": {"CSC100"},
		"CSC100": nil,
	}

	fmt.Println()
	prettyprint(m)
	fmt.Println(solve(m))
	fmt.Println()

	m = map[string][]string{
		"CSC300": {"CSC200"},
		"CSC200": {"CSC100"},
		"CSC100": nil,
	}

	prettyprint(m)
	fmt.Println(solve(m))

}

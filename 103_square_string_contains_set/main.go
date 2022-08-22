package main

import (
	"fmt"
	"strings"
)

func solve(str string, set []rune) string {
	for rng := len(set); rng < len(str); rng++ {
		for i := 0; i <= len(str)-rng; i++ {
			substr := str[i : i+rng]
			exists := true
			for _, rn := range set {
				if strings.IndexRune(substr, rn) == -1 {
					exists = false
				}
			}

			if exists {
				return substr
			}
		}
	}

	return ""
}

func main() {
	fmt.Println(solve("figehaeci", []rune{'a', 'e', 'i'}))
}

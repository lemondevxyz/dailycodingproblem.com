package main

import (
	"fmt"
	"sort"
	"strings"
)

func solve(sen string, words []string) []string {

	// where index is the string index of word
	wm := map[int]string{}
	// slice containing string indexes
	indexes := []int{}

	// unmodified sentence string
	// used to save the index
	orisen := sen

	for _, v := range words {
		if strings.Contains(sen, v) {
			i := strings.Index(orisen, v)
			if i >= 0 {
				// replace the word in the string to make sure no words alike exist.
				// bedbath, bed, bath.
				sen = strings.ReplaceAll(sen, v, "")

				// the index of the word = the word
				wm[i] = v
				// append to indexes slice to sort through later.
				indexes = append(indexes, i)
			}
		}
	}
	// sort em so we know the order of the words
	sort.Ints(indexes)

	// reset the words so we can later return em
	words = []string{}
	// range over sorted indexes
	for _, v := range indexes {
		words = append(words, wm[v])
	}

	return words
}

func main() {
	fmt.Println(solve("thequickbrownfox", []string{"quick", "fox", "brown", "the"}))
	fmt.Println(solve("bedbathandbeyond", []string{"bed", "bath", "bedbath", "and", "beyond"}))
}

package main

import "fmt"

type Flight struct {
	prev string
	next string
}

func count(str string) int {
	res := 0
	for _, v := range str {
		res += int(v)
	}

	return res
}

func solve(fls []Flight, start string) []string {

	index := -1
	for k, v := range fls {
		if v.prev == start {
			index = k
			break
		}
	}

	if index == -1 {
		return nil
	}

	places := []string{}
	places = append(places, fls[index].prev)
	next := fls[index].next

	// avoid repeating yourself
	ignore := map[int]bool{}
	// skip the starting point
	ignore[index] = true

	for i := 0; i < len(fls); i++ {

		indexes := map[int]int{}

		// add flights to compare lexicographically
		for k, v := range fls {
			if _, exist := ignore[k]; exist {
				continue
			}

			if v.prev == next {
				indexes[k] = count(v.prev)
			}
		}

		// no matches
		if len(indexes) <= 0 {
			continue
		}

		// compare lexicographically
		smoll := -1
		for k, v := range indexes {
			if smoll > v || smoll == -1 {
				smoll = v
				index = k
			}
		}

		// add to string list
		if index >= 0 {
			places = append(places, fls[index].prev)
			next = fls[index].next

			ignore[index] = true
		}
	}

	places = append(places, fls[index].next)

	if len(places) != len(fls)+1 {
		return nil
	}

	return places
}

func main() {
	input := [][]Flight{
		{
			{"SFO", "HKO"},
			{"YYZ", "SFO"},
			{"YUL", "YYZ"},
			{"HKO", "ORD"},
		},
		// ['YUL', 'YYZ', 'SFO', 'HKO', 'ORD']
		{
			{"SMO", "COM"},
			{"COM", "YYZ"},
		},
		{
			{"A", "B"},
			{"A", "C"},
			{"B", "C"},
			{"C", "A"},
		},
	}

	str := []string{
		"YUL",
		"COM",
		"A",
	}

	for i := 0; i < len(input); i++ {
		fmt.Println(input[i], str[i], "\n\t", solve(input[i], str[i]))
	}
}

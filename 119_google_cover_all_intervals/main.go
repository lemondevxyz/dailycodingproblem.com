package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func (i Interval) String() string {
	return fmt.Sprintf("[%d, %d]", i.Start, i.End)
}

// how does this work?
// first, we start off by creating a slice (array) of keys
// and a map to hold the keys's values
//
// second we loop through the slice of intervals
// and input the necessary data into the slice and the map
// note: the map is used because it allows for O(1) checking of a key
// existing..
// the necessary data being what Start and End of the interval exist
// within which index of slice.
// for example: [0, 3], [0, 2], [3, 4], [6, 9]
// 3 has indexes 0 and 2
// while 2 has indexes 1
// so, the map would look like: {3: []int{0, 2}, 2: []int{1}}
//
// third, sort through the array because we want to minimize the loop count
// for the integers. because if it wasn't sorted, we'd run into the loop
// running randomly, and the interval having a bigger start than the end
//
// fourth, go through the keys twice because we want to compare the first
// value with the second.
// let's imagine the dailycodingproblem.com input data with is:
// [0, 3], [0, 2], [3, 4], [6, 9]
// the array would be: [0, 2, 3, 4, 6, 9]
// then the loop would compare 0, with the rest of what's ahead
// likewise, 2 would skip over 0 and start with 3 till 9
// and so on.
//
// fifth, within the loop, compare the amount of indexes in each key's
// map, if they equal to the length of the given slice, then move to
// the next step.
//
// sixth, create a map, go through each key's slice of indexes and input
// it into the map. afterwards, if the map's length is equal to the
// slice's length then return the interval that we are in.
// essentially, the amount of keys in the map correspond to the amount
// of *legal* indexes. the data is already safe because we have input
// it from the slice of intervals given to us.
//
// do note: this algorithm doesn't aim to be the fastest.
func solve(slice []Interval) Interval {
	x := map[int][]int{}

	keys := []int{}
	for index, interval := range slice {
		_, ok := x[interval.Start]
		if !ok {
			x[interval.Start] = []int{}
			keys = append(keys, interval.Start)
		}
		x[interval.Start] = append(x[interval.Start], index)

		_, ok = x[interval.End]
		if !ok {
			x[interval.End] = []int{}
			keys = append(keys, interval.End)
		}

		x[interval.End] = append(x[interval.End], index)
	}

	sort.Ints(keys)
	for k1 := range keys {
		for k2 := range keys[k1:] {
			ind1, ind2 := keys[k1], keys[k1+k2]
			if ind1 == ind2 {
				continue
			}

			has1, has2 := x[keys[k1]], x[keys[k1+k2]]

			size := len(has1) + len(has2)
			if size == len(slice) {
				m := map[int]struct{}{}

				for _, v := range has1 {
					m[v] = struct{}{}
				}
				for _, v := range has2 {
					m[v] = struct{}{}
				}

				if len(m) == len(slice) {
					return Interval{ind1, ind2}
				}
			}
		}
	}

	return Interval{}

}

func main() {
	fmt.Println(solve([]Interval{
		{0, 3},
		{2, 6},
		{3, 4},
		{6, 9},
	}))
}

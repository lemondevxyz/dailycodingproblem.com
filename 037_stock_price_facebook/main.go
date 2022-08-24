package main

import "fmt"

func solve(stonks []int) int {

	// BAD
	if len(stonks) == 0 {
		return 0
	}

	maxpro := 0
	start := stonks[0]

	// basically record the profit
	// if it's negative, reset profit and starting point
	// if it's possible profit is bigger than current profit then
	// set current profit to possible profit
	for _, v := range stonks {
		if v == start {
			continue
		}

		profit := v - start
		if profit > maxpro {
			maxpro = profit
		}

		if profit < 0 {
			profit = 0
			start = v
		}
	}

	return start
}

func main() {
	inputs := [][]int{
		{9, 11, 8, 5, 7, 10},
	}

	for _, v := range inputs {
		fmt.Printf("%v: %d\n", v, solve(v))
	}
}

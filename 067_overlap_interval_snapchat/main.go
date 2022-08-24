package main

import "fmt"

type span struct {
	src int
	dst int
}

func solve(s []span) (res []span) {

	skip := map[int]bool{}
	for k := 0; k < len(s); k++ {
		v := s[k]
		for l := k + 1; l < len(s); l++ {
			x := s[l]
			if !(v.src < v.dst && v.dst < x.src && x.src < x.dst) {
				src := v.src
				if v.src > x.src {
					src = x.src
				}

				dst := v.dst
				if v.dst < x.dst {
					dst = x.dst
				}

				skip[k] = true
				skip[l] = true

				res = append(res, span{src, dst})

				break
			}
		}

		if !skip[k] {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	ss := [][]span{
		{
			{1, 3},
			{4, 10},
			{5, 12},
			{20, 25},
		},
		{
			{1, 3},
			{4, 10},
			{5, 8},
			{20, 25},
		},
		{
			{1, 9},
			{7, 10},
			{12, 21},
			{20, 25},
		},
	}
	for _, v := range ss {
		fmt.Println(v, "becomes", solve(v))
	}
}

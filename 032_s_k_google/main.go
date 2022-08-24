package main

import "fmt"

func get_arr(arr []int, offset int) []int {
	ret := arr[offset:]
	return append(ret, arr[:offset]...)
}

func solve(s []int, k int) []int {
	for i := 0; i < len(s); i++ {
		add := 0
		s := get_arr(s, i)
		for j, v := range s {
			add += v
			if add == k {
				return s[:j+1]
			}
		}
	}

	return nil
}

func main() {
	inputs := [][]int{{12, 1, 61, 5, 9, 2}}
	ks := []int{24}
	for k, v := range inputs {
		arr := solve(v, ks[k])
		fmt.Printf("%v: %v\n", v, arr)
	}
}

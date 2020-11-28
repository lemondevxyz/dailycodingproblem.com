package main

import (
	"fmt"
	"sort"
)

type sortedResults [][]int

func (sr sortedResults) Swap(i, j int) {
	sr[i], sr[j] = sr[j], sr[i]
}

func (sr sortedResults) Less(i, j int) bool {
	s1, s2 := len(sr[i]), len(sr[j])

	if s1 == s2 {
		str1, str2 := fmt.Sprintf("%v", sr[i]), fmt.Sprintf("%v", sr[j])

		return str1 < str2
	}

	return s1 < s2
}

func (sr sortedResults) Len() int {
	return len(sr)
}

func filter(arr []int, start int, size int) []int {

	result := []int{}
	for i := start; i < start+size; i++ {
		v := i % len(arr)

		if v < len(arr) {
			result = append(result, arr[v])
		}
	}

	return result
}

// O(n^3), possible bad implementation
func solve(arr []int) [][]int {

	result := sortedResults{}
	result = append(result, []int{})

	exist := map[string]bool{}
	for i := range arr {
		// start from 2
		result = append(result, []int{arr[i]})
		for x := 2; x <= len(arr); x++ {
			val := filter(arr, i, x)
			// pretty output
			sort.Ints(val)
			id := fmt.Sprintf("%v", val)
			if _, ok := exist[id]; !ok {
				result = append(result, val)
				exist[id] = true
			}
		}
	}

	// pretty output
	sort.Sort(result)

	return result
}

func main() {
	arr := []int{1, 2, 3}
	res := solve(arr)
	fmt.Println(arr, ":", res)
}

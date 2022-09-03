package main

import (
	"fmt"
	"sort"
)

type N struct {
	l *N
	v int
	r *N
}

func (n *N) do(arr [][]int, index int) ([][]int, int) {
	original := append([]int{}, arr[index]...)
	if n.l != nil {
		arr[index] = append(append([]int{}, arr[index]...), n.l.v)
		arr, index = n.l.do(arr, index)
	}

	if n.l != nil && n.r != nil {
		arr = append(arr, append([]int{}, original...))
		index++
	}

	if n.r != nil {
		arr[index] = append(append([]int{}, original...), n.r.v)
		arr, index = n.r.do(arr, index)
	}

	return arr, index
}

func sumIt(a []int) (sum int) {
	for _, num := range a {
		sum += num
	}

	return
}

func solve(n *N) int {
	arr, _ := n.do([][]int{{n.v}}, 0)

	m := map[int]int{}
	sums := []int{}

	for _, v := range arr {

		sum := sumIt(v)
		m[sum] = len(v)
		sums = append(sums, sum)
	}
	sort.Ints(sums)

	if len(sums) == 0 {
		return -1
	}

	return m[sums[0]]
}

func main() {
	n := &N{
		&N{
			&N{nil, 3, nil},
			2,
			nil,
		},
		1,
		&N{
			&N{nil, 5, &N{nil, -20, nil}},
			12,
			nil,
		},
	}

	fmt.Println(solve(n))
}

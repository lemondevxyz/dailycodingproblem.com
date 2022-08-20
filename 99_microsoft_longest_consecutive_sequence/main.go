package main

import (
	"fmt"
)

type Value struct {
	Number   int
	Left     *Value
	LeftVal  int
	Right    *Value
	RightVal int
}

func (v *Value) IncreaseLeftf() {
	if v.Left != nil {
		v.Left.IncreaseLeftf()
	} else {
		v.Number++
	}
}

func (v *Value) IncreaseRightf() {
	if v.Left != nil {
		v.Left.IncreaseLeftf()
	} else {
		v.Number++
	}
}

func solve(arr []int) int {

	m := map[int]*int{}
	slice := []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]

		num := 1

		vminus := v - 1
		vplus := v + 1

		val, ok := m[v]
		if !ok {
			m[vplus] = &num
			m[v] = &num
			m[vminus] = &num
		} else {
			num = *val + 1
			*val = num
			m[vplus] = val
			m[v] = val
			m[vminus] = val
		}

		if len(slice) == 0 {
			slice = append(slice, num)
		} else {
			have := slice[0]
			if num > have {
				slice[0] = num
			}
		}
	}

	return slice[0]
}

func main() {
	fmt.Println(solve([]int{100, 4, 200, 1, 3, 2}))
}

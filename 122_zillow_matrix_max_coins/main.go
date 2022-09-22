package main

import "fmt"

type matrix [][]int

func (m matrix) String() string {
	str := ""
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			str += fmt.Sprintf("%d ", m[i][j])
		}

		str += "\n"
	}

	return str[:len(str)-1]
}

func (m matrix) Bounds() int {
	return len(m) * len(m[0])
}

func solve(m matrix, paths [][]int, index int) ([][]int, int) {
	slice := paths[index]
	place := slice[len(slice)-1]

	maxX := len(m[0])
	down := place + maxX

	right := place + 1

	max := m.Bounds()
	original := append([]int{}, slice...)
	if down < max {
		paths[index] = append(paths[index], down)

		paths, index = solve(m, paths, index)
	}

	if down < max && right < max && right%maxX >= place%maxX {
		paths = append(paths, original)
		index++
	}

	if right < max && right%maxX >= place%maxX {
		paths[index] = append(paths[index], right)

		paths, index = solve(m, paths, index)
	}

	return paths, index
}

func sumSlice(a []int) (res int) {
	for _, v := range a {
		res += v
	}

	return
}

func main() {
	m := matrix{
		[]int{0, 3, 1, 1},
		[]int{2, 0, 0, 4},
		[]int{1, 5, 3, 1}}

	paths, _ := solve(m, [][]int{{0}}, 0)
	solve(m, paths, 0)

	sum := 0
	slice := []int{}

	max := len(m[0])
	for _, arr := range paths {
		values := []int{}

		for _, v := range arr {
			y := v / max
			x := v % max

			values = append(values, m[y][x])
		}

		sliceSum := sumSlice(values)
		if sliceSum > sum {
			slice = values
			sum = sliceSum
		}
	}

	fmt.Println("maximum sum:", sum)
	fmt.Println("steps:", slice)
}

package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Stolen from 110_facebook_node_all_paths
func (n *Node) Do(arr [][]int, index int) ([][]int, int) {
	original := append([]int{}, arr[index]...)
	if n.Left != nil {
		arr[index] = append(arr[index], n.Left.Val)
		arr, index = n.Left.Do(arr, index)
	}

	if n.Left != nil && n.Right != nil {
		arr = append(arr, []int{n.Val})
		index++
	}

	if n.Right != nil {
		arr[len(arr)-1] = append(append([]int{}, original...), n.Right.Val)
		arr, index = n.Right.Do(arr, index)
	}

	return arr, index
}

func (n *Node) string(lvl int) string {
	s := fmt.Sprintf("%s%d\n", strings.Repeat(" ", lvl), n.Val)

	if n.Left != nil {
		s += "l" + n.Left.string(lvl+1)
	}

	if n.Right != nil {
		s += "r" + n.Right.string(lvl+1)
	}

	return s
}

func retArray(s []int, want int) []int {
	d := []int{}
	for index, have := range s {
		if have == want {
			for _, num := range s[:index] {
				d = append(d, num)
			}

			return d
		}
	}

	return d
}

func solve(n *Node, l, r int) int {
	if n == nil {
		return -1
	}

	arr, _ := n.Do([][]int{{n.Val}}, 0)

	var lArr, rArr []int

	for _, v := range arr {
		if len(lArr) == 0 {
			lArr = retArray(v, l)
		}

		if len(rArr) == 0 {
			rArr = retArray(v, r)
		}
	}

	var ind int
	for i := 0; i < len(lArr); i++ {
		for j := 0; j < len(rArr); j++ {

			if lArr[i] == rArr[j] {
				if i > ind {
					ind = i
				}
			}
		}
	}

	return lArr[ind]
}

func main() {
	n := &Node{
		Val:   1,
		Left:  &Node{2, &Node{3, nil, &Node{7, nil, &Node{8, nil, nil}}}, nil},
		Right: &Node{4, &Node{5, nil, &Node{6, nil, nil}}, nil},
	}

	fmt.Println("The binary tree:")
	fmt.Println(n.string(0))

	fmt.Println()

	l, r := 2, 3
	fmt.Printf("The lowest common ancestor between %d and %d\n", l, r)
	fmt.Println(solve(n, l, r))
}

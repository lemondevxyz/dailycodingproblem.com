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

func solve(n *Node, want int) (*int, *int) {
	arr, _ := n.Do([][]int{{n.Val}}, 0)
	for k, slice1 := range arr {
		for _, slice2 := range arr[k+1:] {
			for i1, num1 := range slice1 {
				for i2, num2 := range slice2 {
					if i1 == i2 && num1 == num2 {
						continue
					}

					plus := num1 + num2
					if plus == want {
						return &num1, &num2
					}
				}
			}
		}
	}

	return nil, nil
}

func main() {
	n := &Node{
		10,
		&Node{5, nil, nil},
		&Node{
			15,
			&Node{11, nil, nil},
			&Node{15, nil, nil},
		},
	}

	fmt.Print(n.string(0))
	n1, n2 := solve(n, 20)
	fmt.Println("result", *n1, *n2)
}

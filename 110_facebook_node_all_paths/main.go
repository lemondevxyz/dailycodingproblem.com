package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Do(arr [][]int, index int) ([][]int, int) {
	original := append([]int{}, arr[index]...)
	if n.Left != nil {
		arr[index] = append(append([]int{}, arr[index]...), n.Left.Val)
		arr, index = n.Left.Do(arr, index)
	}

	if n.Left != nil && n.Right != nil {
		arr = append(arr, append([]int{}, original...))
		index++
	}

	if n.Right != nil {
		arr[index] = append(append([]int{}, original...), n.Right.Val)
		arr, index = n.Right.Do(arr, index)
	}

	return arr, index
}

func main() {
	n := &Node{
		1,
		&Node{
			2,
			&Node{
				4,
				&Node{
					5, nil, nil,
				},
				nil,
			},
			&Node{
				3, nil, nil,
			},
		},
		&Node{
			6,
			nil,
			&Node{
				7,
				&Node{
					8,
					&Node{
						10, nil, nil,
					},
					&Node{
						9, nil, nil,
					},
				},
				nil,
			},
		},
	}

	arr, _ := n.Do([][]int{{n.Val}}, 0)
	fmt.Println(arr)
}

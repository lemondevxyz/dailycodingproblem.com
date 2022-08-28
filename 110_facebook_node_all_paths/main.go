package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Do(arr [][]int, index int) [][]int {
	addIndex := false
	if n.Left != nil && n.Right != nil {
		addIndex = true
	}

	if n.Left != nil {
		arr[index] = append(arr[index], n.Left.Val)
		n.Left.Do(arr, index)
	}

	if addIndex {
		slice := arr[index]
		if len(slice) != 0 {
			slice = slice[:len(slice)-1]
		}

		arr = append(arr, append([]int{}, slice...))
		index++
	}

	if n.Right != nil {
		arr[index] = append(arr[index], n.Right.Val)
		arr = n.Right.Do(arr, index)
	}

	return arr
}

func main() {
	n := &Node{1, &Node{2, nil, nil}, &Node{3, &Node{4, nil, nil}, &Node{5, nil, nil}}}

	arr := n.Do([][]int{{n.Val}}, 0)
	fmt.Println(arr)
}

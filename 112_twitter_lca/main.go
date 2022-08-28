package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Stolen from 110_facebook_node_all_paths
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

func max(arr [][]int) (int, int) {
	index := -1
	size := 0
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		if len(v) > size || index == -1 {
			fmt.Println(v)
			size = len(v)
			index = i
		}
	}

	return size, index
}

func subseq(arr [][]int, index int) [][]int {
	if index == 0 {
		if len(arr) > 1 {
			return arr[1:]
		}

		return nil
	} else if index == len(arr)-1 {
		return arr[:len(arr)-1]
	}

	return append(arr[:index], arr[index:]...)
}

func solve(n *Node) int {
	if n == nil {
		return -1
	}

	arr := n.Do([][]int{{n.Val}}, 0)
	fmt.Println(arr)
	fmt.Println(max(arr))

	return 0
}

func main() {
	n := &Node{
		Val:   1,
		Left:  &Node{2, &Node{3, nil, &Node{7, nil, &Node{8, nil, nil}}}, nil},
		Right: &Node{4, &Node{5, nil, &Node{6, nil, nil}}, nil},
	}
	solve(n)
}

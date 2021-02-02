package main

import "fmt"

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

const horz = "|"
const vert = "─"

func (n *Node) String() string {
	return n.Val
}

func (n *Node) Invert() {

	n.Left, n.Right = n.Right, n.Left
	if n.Left != nil {
		n.Left.Invert()
	}

	if n.Right != nil {
		n.Right.Invert()
	}
}

/*
func (n *Node) Array(arr [][]*Node, level int) [][]*Node {

	arr = append(arr, []*Node{})
	arr[level] = append(arr[level], n)
	level++

	if n.Left != nil || n.Right != nil {
		arr = append(arr, []*Node{})
	}

	if n.Left != nil {
		arr[level] = append(arr[level], n.Left)

		n.Left.Array(arr, level+1)
	}

	if n.Right != nil {
		arr[level] = append(arr[level], n.Right)

		n.Right.Array(arr, level+1)
	}

	return arr
}
*/

func main() {
	n := Node{
		Val: "a",
		Left: &Node{
			Val: "b",
			Left: &Node{
				Val: "d",
			},
			Right: &Node{
				Val: "e",
			},
		},
		Right: &Node{
			Val: "c",
			Left: &Node{
				Val: "f",
			},
		},
	}

	fmt.Println("left:", n.Left)
	fmt.Println("left left:", n.Left.Left)
	fmt.Println("left right:", n.Left.Right)
	fmt.Println("right:", n.Right)
	fmt.Println("right left:", n.Right.Left)
	n.Invert()
	fmt.Println("inverting")
	fmt.Println("left:", n.Left)
	fmt.Println("left right:", n.Left.Right)
	fmt.Println("right:", n.Right)
	fmt.Println("right left:", n.Right.Left)
	fmt.Println("right left:", n.Right.Right)
}

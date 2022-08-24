package main

import "fmt"

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

func (n Node) DeepestNode(size int) (int, *Node) {
	var lefi int
	var lefn *Node
	if n.Left != nil {
		lefi, lefn = n.Left.DeepestNode(size + 1)
	}

	var rigi int
	var rign *Node
	if n.Right != nil {
		rigi, rign = n.Right.DeepestNode(size + 1)
	}

	if rigi >= lefi && rign != nil {
		return rigi, rign
	} else if lefi >= rigi && lefn != nil {
		return lefi, lefn
	}

	return size, &n
}

func main() {
	n := Node{
		Val: "a",
		Left: &Node{
			Val: "b",
			Left: &Node{
				Val: "d",
			},
		},
		Right: &Node{
			Val: "c",
		},
	}

	x, z := n.DeepestNode(0)
	fmt.Printf("Deepest node is '%s', with level %d.\n", z.Val, x)
}

package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n Node) Valid(root, i int) bool {
	oklef, okrig := false, false
	if n.Left != nil {
		if n.Left.Val <= root {
			oklef = n.Left.Valid(root, i+1)
		}
	} else {
		oklef = true
	}

	if n.Right != nil {
		if n.Right.Val >= root {
			okrig = n.Right.Valid(root, i+1)
		}
	} else {
		okrig = true
	}

	if oklef && okrig {
		return true
	}

	return false
}

func main() {
	tree := Node{
		Val: 5,
		Left: &Node{
			Val: 4,
			Left: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 6,
			Left: &Node{
				Val: 2,
			},
			Right: &Node{
				Val: 9,
			},
		},
	}

	fmt.Printf("root val is equal to %d\n", tree.Val)
	fmt.Printf("first one should be okay: %t\n", tree.Valid(tree.Val, 0))
	tree.Right.Left.Val = 8
	fmt.Printf("tree->right->left = %d: %t\n", tree.Right.Left.Val, tree.Valid(tree.Val, 0))
	tree.Right.Left.Val = 1
	tree.Left.Right = &Node{Val: 4}
	fmt.Printf("tree->right->left = %d\n", tree.Right.Left.Val)
	fmt.Printf("tree->left->right = %d: %t\n", tree.Left.Right.Val, tree.Valid(tree.Val, 0))
}

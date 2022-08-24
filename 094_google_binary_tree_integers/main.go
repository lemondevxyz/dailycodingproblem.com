package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Rightie(i int) (*Node, int) {
	r := n.Right

	if r != nil {
		r, rn := n.Right.Rightie(i + 1)
		ln := 0
		l := n.Left
		if l != nil {
			l, ln = n.Left.Rightie(i + 1)
		}

		if rn > ln {
			return r, rn
		} else {
			return l, ln
		}
	}

	return n, i
}

func (n *Node) Leftie(i int) (*Node, int) {
	l := n.Left

	if l != nil {
		l, ln := n.Left.Leftie(i + 1)
		rn := 0
		r := n.Right
		if r != nil {
			r, rn = n.Right.Leftie(i + 1)
		}

		if rn > ln {
			return r, rn
		} else {
			return l, ln
		}
	}

	return n, i
}

func main() {
	root := &Node{
		1,
		&Node{2, &Node{3, nil, nil}, &Node{4, nil, &Node{5, nil, &Node{6, nil, &Node{7, nil, &Node{8, nil, &Node{9, nil, nil}}}}}}}, &Node{10, nil, nil}}

	l, ln := root.Leftie(1)
	r, rn := root.Rightie(1)
	fmt.Println(l, ln)
	fmt.Println(r, rn)
}

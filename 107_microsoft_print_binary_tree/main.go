package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Left  *Node
	Val   int
	Right *Node
}

func (n *Node) String() string {
	str := strconv.Itoa(n.Val)
	if n.Left != nil {
		str += ", " + n.Left.String()
	}

	if n.Right != nil {
		str += ", " + n.Right.String()
	}

	return str
}

func main() {
	n := &Node{&Node{nil, 2, nil}, 1, &Node{&Node{nil, 4, nil}, 3, &Node{nil, 5, nil}}}

	fmt.Println(n)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Children []*Node
	Size     int
}

// if true then stop
func (n *Node) Dive(level int, param func(int, *Node) bool) {
	if n.Children == nil || len(n.Children) == 0 {
		return
	}

	for _, v := range n.Children {
		if param(level+1, v) {
			break
		} else {
			v.Dive(level+1, param)
		}
	}

}

func solve(god *Node) *Node {
	curr := 0

	var pointcurr *Node
	var pointprev *Node

	god.Dive(0, func(level int, n *Node) bool {
		if n.Size > curr {
			curr = n.Size

			pointprev = pointcurr
			pointcurr = n
		}

		return false
	})

	return pointprev
}

func (n *Node) String() string {
	return strconv.Itoa(n.Size)
}

func (n *Node) Tree() string {

	str := "- " + n.String() + "\n"

	n.Dive(1, func(level int, n *Node) bool {
		str += "-" + strings.Repeat(" ", (level*2)-1) + n.String() + "\n"

		return false
	})

	return str
}

func main() {
	god := &Node{
		Size: 15,
	}
	god.Children = []*Node{
		{
			Size: 65,
		},
		{
			Size: 90,
		},
		{
			Size: 10,
		},
		{
			Size: 90,
			Children: []*Node{
				// we want this one
				{
					Size: 100,
				},
				{
					Size: 165,
				},
			},
		},
	}

	secondlargest := solve(god)
	fmt.Println("following tree:")
	fmt.Println(god.Tree())
	fmt.Println("second largest is", secondlargest)
}

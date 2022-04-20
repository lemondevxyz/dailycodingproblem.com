package main

import "fmt"

type Node struct {
	Val    int
	Parent *Node
	L      *Node
	R      *Node
}

func (n *Node) Do(fn func(*Node, int), lvl int) {
	if n.L != nil {
		n.L.Parent = n
		fn(n.L, lvl+1)
	}
	if n.R != nil {
		n.R.Parent = n
		fn(n.R, lvl+1)
	}
	if n.L != nil {
		n.L.Do(fn, lvl+1)
	}
	if n.R != nil {
		n.R.Do(fn, lvl+1)
	}
}

func main() {
	n := &Node{
		1,
		nil,
		&Node{
			2,
			nil,
			&Node{3, nil, &Node{4, nil, nil, &Node{5, nil, nil, nil}}, nil},
			nil,
		},
		&Node{
			6,
			nil,
			nil,
			&Node{7, nil, nil, &Node{8, nil, nil, nil}},
		},
	}

	maxlvl := 0
	var maxnode *Node
	n.Do(func(n *Node, lvl int) {
		if lvl > maxlvl {
			maxlvl = lvl
			maxnode = n
		}
	}, 0)

	parent := maxnode
	path := ""

	for parent != nil {
		oldnode := parent
		parent = parent.Parent
		if parent != nil {
			if oldnode == parent.R {
				path += ">- R "
			} else if oldnode == parent.L {
				path += ">- L "
			}
		}
	}

	newpath := ""
	for k := len(path) - 1; k >= 0; k-- {
		newpath += string(path[k])
	}
	fmt.Println("Root ->"+newpath[:len(newpath)-2]+"=", maxnode.Val)
}

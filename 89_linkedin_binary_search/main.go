package main

import (
	"fmt"
	"log"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Do(fn func(i int, n *Node)) {
	l, r := n.Left, n.Right

	if l != nil {
		fn(1, l)
		l.Do(fn)
	}

	if r != nil {
		fn(2, r)
		r.Do(fn)
	}
}

func solve(root *Node) (invalid *bool) {
	want := root.Val

	root.Do(func(i int, n *Node) {
		if invalid != nil {
			return
		}
		have := n.Val
		if i == 1 {
			if have > want {
				invalid = new(bool)
				*invalid = true
			}
		} else if i == 2 {
			if have < want {
				invalid = new(bool)
				*invalid = true
			}
		}
	})

	if invalid == nil {
		invalid = new(bool)
		*invalid = false
	}

	return
}

func main() {
	fail := &Node{
		10,
		&Node{
			4,
			&Node{
				11,
				nil,
				nil,
			},
			&Node{
				12,
				nil,
				nil,
			},
		},
		nil,
	}
	success := &Node{
		10,
		&Node{
			4,
			&Node{
				10,
				nil,
				nil,
			},
			&Node{
				12,
				nil,
				nil,
			},
		},
		&Node{
			13,
			nil,
			nil,
		},
	}

	if !(*solve(fail)) {
		log.Fatalf("Not working: failure succeed")
	}

	if *solve(success) {
		log.Fatalf("Not working: success failed")
	}

	fmt.Println("worke")
}

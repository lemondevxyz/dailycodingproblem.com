package main

import (
	"fmt"
	"os"
)

type Node struct {
	Parent   *Node
	Locked   bool
	Children []*Node
}

// callback is function that gets called into all children
// if callback returns true that means stop
// else continue diving
func (n *Node) Dive(callback func(n *Node) bool) {
	for _, v := range n.Children {
		n = v
		if callback(n) {
			break
		} else {
			n.Dive(callback)
		}
	}
}

// same as dive but opposite, meaning we search for every parent.
func (n *Node) Climb(callback func(n *Node) bool) {
	if n.Parent != nil {
		p := n.Parent
		if callback(p) {
			return
		} else {
			p.Climb(callback)
		}
	}
}

func (n *Node) canLock() bool {
	ok := true

	if n.Locked {
		return false
	}

	callback := func(n *Node) bool {
		if n.Locked {
			ok = false
			return true
		}

		return false
	}

	n.Dive(callback)
	n.Climb(callback)

	return ok
}

func (n *Node) Lock() (success bool) {
	if n.canLock() {
		n.Locked = true
		n.Dive(func(n *Node) bool {
			n.Locked = true

			return false
		})

		success = true
	}

	return
}

func (n *Node) Unlock() (success bool) {
	if n.canLock() {
		n.Locked = false
		n.Dive(func(n *Node) bool {
			n.Locked = false

			return false
		})

		success = true
	}

	return
}

func (n *Node) IsLocked() bool { return n.Locked }
func (n *Node) Append(a *Node) *Node {

	a.Parent = n
	if n.Children == nil {
		n.Children = []*Node{}
	}

	n.Children = append(n.Children, a)

	return a
}

func main() {
	// creator of all
	god := Node{}

	jesus := god.Append(&Node{})
	jesus.Append(&Node{}).Append(&Node{})

	// case 1
	if god.Lock() == false {
		fmt.Println("bad code, case 1")
		os.Exit(1)
	}

	// case 2
	if god.Unlock() == true {
		fmt.Println("bad code, case 2")
		os.Exit(1)
	}

	// case 3
	god.Locked = false
	if god.Lock() == true {
		fmt.Println("bad code, case 3")
		os.Exit(1)
	}

	fmt.Println("success")
}

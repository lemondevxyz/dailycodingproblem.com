package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Node struct {
	l *Node
	v int
	r *Node
}

func (n *Node) string(lvl int) string {
	str := fmt.Sprintf("%s%d\n", strings.Repeat(" ", lvl), n.v)
	if n.l != nil {
		str += "l " + n.l.string(lvl+1)
	}

	if n.r != nil {
		str += "r " + n.r.string(lvl+1)
	}

	return str
}

func (n *Node) generate(level, max int) {
	next := level + 1
	if next == max {
		return
	}

	if rand.Intn(2) == 1 {
		n.l = &Node{nil, 1, nil}
		n.l.generate(next, max)
	} else {
		n.r = &Node{nil, 1, nil}
		n.r.generate(next, max)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	n := &Node{nil, 0, nil}
	n.generate(0, 15)
	fmt.Println(n.string(0))
}

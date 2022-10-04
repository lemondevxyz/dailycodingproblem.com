package main

import (
	"fmt"
	"math"
)

type Node struct {
	v int
	r *Node
}

func (n *Node) String() string {
	if n.r == nil {
		return fmt.Sprintf("%d", n.v)
	}

	return fmt.Sprintf("%d -> %s", n.v, n.r.String())
}

func (n *Node) Number(pow int) Number {
	val := n.v * int(math.Pow(10, float64(pow)))
	if n.r == nil {
		return Number(val)
	}

	return Number(val) + n.r.Number(pow+1)
}

type Number int

var m = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func (n Number) String() string {
	return fmt.Sprintf("%d", n)
}

func (n Number) Node() *Node {
	var node *Node
	str := n.String()

	for i := range str {
		node = &Node{m[str[i]], node}
	}

	return node
}

func solve(a, b *Node) *Node {
	fmt.Println(b.Number(0))
	return (a.Number(0) + b.Number(0)).Node()
}

func main() {
	a, b := &Node{9, &Node{9, nil}}, &Node{5, &Node{2, nil}}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(solve(a, b))
}

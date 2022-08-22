package main

import "fmt"

type Node struct {
	Val   int
	Right *Node
}

func (n *Node) String() string {
	if n.Right != nil {
		return fmt.Sprintf("%d -> %s", n.Val, n.Right.String())
	}

	return fmt.Sprintf("%d", n.Val)
}

func NewDoubleLinkedList(a ...int) *Node {
	if len(a) == 0 {
		return nil
	}

	node := &Node{Val: a[0]}
	start := node
	for _, v := range a[1:] {
		node.Right = &Node{Val: v}
		node = node.Right
	}

	return start
}

func (n *Node) Last() int {
	if n.Right != nil {
		return n.Right.Last()
	}

	return n.Val
}

func palindrome(n *Node) bool {
	return n.Val == n.Last()
}

func main() {
	list := NewDoubleLinkedList(1, 4, 3, 2, 1)
	fmt.Println("list:", list)
	fmt.Println("palindrome:", palindrome(list))
	fmt.Println()
	list = NewDoubleLinkedList(1, 4)
	fmt.Println("list:", list)
	fmt.Println("palindrome:", palindrome(list))
	fmt.Println()
}

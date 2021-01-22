package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type Llist struct {
	head *Node
	size int
}

// strategy
// set n as head initially
// then change to n
// also keep foot as reverse of head
func (l *Llist) Reverse() *Llist {
	n := l.head
	if n == nil {
		return &Llist{
			size: 0,
		}
	}

	a := []int{}
	// surely there is a better way to do this
	// but my methods lack simplicity
	for {
		if n == nil {
			break
		} else {
			a = append([]int{
				n.val,
			}, a...)
		}

		n = n.next
	}

	return newList(a)
}

func (l *Llist) String() string {

	n := l.head
	str := "["
	if l.head != nil {
		str += fmt.Sprintf("%d", l.head.val)
	} else {
		return "[]"
	}

	for {
		n = n.next
		if n == nil {
			break
		}

		str += fmt.Sprintf(" %d", n.val)
	}

	return str + "]"
}

func newList(arr []int) *Llist {
	if len(arr) == 0 || arr == nil {
		return nil
	}

	head := &Node{
		val: arr[0],
	}

	var last *Node

	for _, v := range arr[1:] {
		n := &Node{
			val: v,
		}
		if head.next == nil {
			head.next = n
		}

		if last == nil {
			last = n
		} else {
			last.next = n
			last = n
		}
	}

	l := &Llist{
		head: head,
		size: len(arr),
	}

	return l
}

func main() {
	l := newList([]int{1, 2, 3, 4, 5})

	fmt.Println(l, "becomes", l.Reverse())
}

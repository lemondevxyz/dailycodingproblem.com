package main

import (
	"fmt"
	"sort"
)

type Node struct {
	val  int
	next *Node
}

type Llist struct {
	head *Node
	size int
}

// so i wasn't sure what was asked, i ended up implementing sort.Interface to sort through the linked list
// also implemented the Append from 63_linked_list_reverse_google

// sort.Interface
func (l *Llist) Len() int {
	return l.size
}

func (l *Llist) Less(i, j int) bool {
	x, _ := l.Get(i)
	y, _ := l.Get(j)

	return x < y
}

func (l *Llist) Swap(i, j int) {
	x, _ := l.Get(i)
	y, _ := l.Get(j)

	l.Set(i, y)
	l.Set(j, x)
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

func (l *Llist) Get(x int) (res int, ok bool) {
	h := l.head
	if h == nil {
		return
	}

	if x >= l.size {
		return
	}

	n := h

	for i := 0; i < l.size; i++ {
		if i == x {
			res, ok = n.val, true
			break
		}

		n = n.next
	}

	return
}

func (l *Llist) Set(x int, val int) (ok bool) {
	h := l.head
	if h == nil {
		return
	}

	if x >= l.size {
		return
	}

	n := h

	for i := 0; i < l.size; i++ {
		if i == x {
			ok = true
			n.val = val
			break
		}

		n = n.next
	}

	return
}

func (l *Llist) Append(dst *Node) {

	if dst == nil {
		return
	}

	n := l.head
	if l.head == nil {
		l.head = n
		return
	}

	for {
		p := n.next
		if p == nil {
			break
		} else {
			n = p
		}
	}

	x := dst
	i := 0

	for {
		i++
		p := x.next
		if p == nil {
			break
		}
		x = x.next
	}

	n.next = dst
	l.size += i

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
	l1 := newList([]int{1, 2, 5, 9, 10})
	l2 := newList([]int{6, 7, 8, 14})
	l3 := newList([]int{12, 15, 19})

	fmt.Println(l1, l2, l3)
	fmt.Println("becomes")
	l2.Append(l3.head)
	l1.Append(l2.head)
	fmt.Println(l1)
	fmt.Println("after sort")
	sort.Sort(l1)
	fmt.Println(l1)
}

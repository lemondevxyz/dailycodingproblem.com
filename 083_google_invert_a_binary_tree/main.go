package main

import (
	"fmt"
	"io"
	"os"
)

type Tree struct {
	Val   byte
	Left  *Tree
	Right *Tree
}

func (t *Tree) Do(fn func(t *Tree)) {
	fn(t)

	if t.Left != nil {
		t.Left.Do(fn)
	}
	if t.Right != nil {
		t.Right.Do(fn)
	}
}

func (t *Tree) Print(w io.Writer, ns int, ch rune) {
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%c\n", ch, t.Val)
	if t.Left != nil {
		t.Left.Print(w, ns+2, 'L')
	}

	if t.Right != nil {
		t.Right.Print(w, ns+2, 'R')
	}
}

func main() {
	t := &Tree{
		Val: 'a',
		Left: &Tree{
			Val: 'b',
			Left: &Tree{
				Val: 'd',
			},
			Right: &Tree{
				Val: 'e',
			},
		},
		Right: &Tree{
			Val: 'c',
			Left: &Tree{
				Val: 'f',
			},
		},
	}

	fmt.Println(`= refers to the root
R refers to the right node
L refers to the left node

before:`)
	t.Print(os.Stdout, 0, '=')

	t.Do(func(t *Tree) {
		t.Left, t.Right = t.Right, t.Left
	})

	fmt.Printf("\nafter:\n")
	t.Print(os.Stdout, 0, '=')

}

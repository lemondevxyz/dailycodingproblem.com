package main

import "fmt"

type fifo []string

func (f *fifo) enqueue(str string) {
	if f == nil {
		return
	}

	*f = append(*f, str)
}

func (f *fifo) dequeue() {
	if len(*f) == 0 || f == nil {
		return
	}

	if len(*f) > 1 {
		slice := *f
		*f = slice[1:]
	} else {
		*f = []string{}
	}
}

func main() {
	f := &fifo{}

	f.enqueue("first item")
	f.enqueue("middle item")
	f.enqueue("last item")

	fmt.Println(f)
	f.dequeue()
	fmt.Println(f)
	f.dequeue()
	fmt.Println(f)
	f.dequeue()
	fmt.Println(f)
}

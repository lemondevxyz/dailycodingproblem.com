package main

import "fmt"

type ele struct {
	name string
	next *ele
}

type singleList struct {
	length int
	head   *ele
}

func initList() *singleList {
	return &singleList{}
}

func (s *singleList) Prepend(name string) {
	el := &ele{
		name: name,
	}

	if s.head != nil {
		s.head = el
	} else {
		el.next = s.head
		s.head = el
	}

	s.length++
}

func (s *singleList) Append(name string) {
	ele := &ele{
		name: name,
	}

	if s.head == nil {
		s.head = ele
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}

		current.next = ele
	}

	s.length++
}

func (s *singleList) Remove(k int) {
	if s.head == nil {
		return
	}

	if k == 0 {
		s.head = s.head.next
		return
	}

	prev := s.head
	current := prev
	for current.next != nil {
		if k == 0 {
			break
		}
		prev = current
		current = current.next
		k--
	}

	if prev != nil {
		// where current is the element we want to skip
		prev.next = current.next
	} else {
		s.head = nil
	}
	s.length--

}

func (s *singleList) Front() string {
	if s.head == nil {
		return ""
	}

	return s.head.name
}

func (s *singleList) Back() string {
	if s.head == nil {
		return ""
	}

	prev := s.head
	current := prev
	for current != nil {
		prev = current
		current = current.next
	}

	return prev.name
}

func (s *singleList) String() string {
	if s.head == nil {
		return "[]"
	}

	str := ""

	var prev *ele

	current := s.head
	for current != nil {
		prev = current
		current = current.next

		str += prev.name + ", "
	}
	// space and ,
	str = str[:len(str)-2]

	return "[" + str + "]"
}

// https://golangbyexample.com/singly-linked-list-in-golang/
// i was pretty lost until i stumbled upon this
func main() {
	l := initList()
	l.Append("A")
	l.Append("B")
	l.Append("C")
	l.Append("D")
	l.Append("E")
	l.Append("F")

	fmt.Println(l)
	l.Remove(0)
	fmt.Println(l)
	l.Remove(2) // B C D E F
	// remove D
	fmt.Println(l)
}

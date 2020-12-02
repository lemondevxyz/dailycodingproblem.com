// i think this is o(1)
// not too sure tho
package main

import "fmt"

type stack struct {
	data []string
}

func (s *stack) push(str string) {
	s.data = append(s.data, str)
}

func (s *stack) pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *stack) max() int {
	return len(s.data)
}

func main() {
	s := stack{}
	s.push("ayy")
	s.push("ababa")
	s.push("agagds")

	s.pop()
	fmt.Println(s.data)
}

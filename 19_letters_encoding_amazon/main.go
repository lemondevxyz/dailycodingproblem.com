package main

import (
	"bytes"
	"fmt"
)

type runlength string

func isNumber(char byte) bool {
	return char >= 48 && char <= 57
}

// O(2n)
func (r *runlength) Decode() {

	body := []byte(*r)
	max := len(body)

	numbers := []int{}

	// first repeat the letters
	for i := 0; i < max; i++ {

		numbers = append(numbers, i)
		char := body[i]

		if isNumber(char) {
			i++

			// how many times to repeat the letter
			repeat := int(char-48) - 1
			// cause we already have a letter
			v := body[i]

			oldbyt := []byte{}

			// [4, 3, 2, 1, 5, 6]
			oldbyt = append(oldbyt, body...)

			// repeat the letter
			body = bytes.Repeat([]byte{v}, repeat)

			if i == 0 {
				body = append(body, oldbyt...)
			} else if i+1 == len(oldbyt) {
				body = append(oldbyt, body...)
			} else {
				if i >= len(oldbyt) {
					i = len(oldbyt)
				}

				body = append(body, oldbyt[i:]...)

				if i < len(oldbyt) {
					body = append(oldbyt[:i], body...)
				}
				//fmt.Println(string(body))
			}

			i += repeat
		}

		max = len(body)
	}

	result := []byte{}
	// second remove the numbers
	for _, v := range body {
		if !isNumber(v) {
			result = append(result, v)
		}
	}

	str := runlength(result)
	*r = str

}

func (r *runlength) Encode() {

	body := []byte(*r)
	result := []byte{}

	start := -1
	for i := 0; i < len(body); i++ {
		if i > 0 {
			prev := i - 1
			want, have := body[i], body[prev]
			if want == have {
				if start == -1 {
					start = prev
				}

				if i+1 == len(body) {
					have = byte(0)
					i++
				}
			}

			if want != have {
				if start >= 0 {
					repeat := len(body[start:i])

					start := i - repeat
					if start > len(body) {
						start = len(body)
					}

					result = append(result, []byte{
						byte(repeat + 48),
						body[prev],
					}...)

				} else {
					// special case for characters occurring once
					result = append(result, []byte{
						byte(49),
						body[prev],
					}...)
				}

				start = -1
			}
		}
	}

	str := string(result)
	*r = runlength(str)
}

func (r *runlength) String() string {
	return string(*r)
}

func create(str string) *runlength {
	v := runlength(str)
	return &v
}

func main() {
	r := create("4A3B2C1D2A")
	r.Decode()

	fmt.Println(r)
	r.Encode()
	fmt.Println(r)
}

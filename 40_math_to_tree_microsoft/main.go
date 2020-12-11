package main

import (
	"fmt"
	"strconv"
	"strings"
)

// lowkey this was pretty hard
type Op uint8

const (
	Add Op = iota + 1
	Sub
	Mul
	Div
)

var Ops = []Op{
	Add,
	Sub,
	Mul,
	Div,
}

func (o Op) Calculate(x, y int) int {
	switch o {
	case Add:
		return x + y
	case Sub:
		return x - y
	case Mul:
		return x * y
	case Div:
		return x / y
	}

	return x
}

func (o Op) String() string {
	switch o {
	case Add:
		return "+"
	case Sub:
		return "-"
	case Mul:
		return "*"
	case Div:
		return "/"
	}

	return ""
}

func get_op(str string) Op {
	for _, v := range Ops {
		if str == v.String() {
			return v
		}
	}

	return Op(0)
}

func solve(str string) int {
	spl := strings.Split(str, "\n")
	// cause we have slashes every other line
	numbers := [][2]int{}
	exist := map[int]bool{}
	// start from down to top
	// also skip any slashes
	for i := len(spl) - 1; i >= 0; i -= 2 {

		substr := spl[i]
		// well, remove all unneccessary space
		substr = strings.Join(strings.Fields(substr), " ")
		fmt.Println(substr)

		v := strings.Split(substr, " ")

		// go through each entry in the string
		for k := range v {
			bcur := v[k]
			ncur, e1 := strconv.Atoi(bcur)
			// if it's not number
			if e1 == nil {
				// then don't peek to next number :)
				if k+1 < len(v) {
					bnex := v[k+1]
					nnex, e2 := strconv.Atoi(bnex)
					if e1 == nil || e2 == nil {
						_, ok1 := exist[k]
						_, ok2 := exist[k+1]
						if !ok1 && !ok2 {
							exist[k] = true
							exist[k+1] = true

							numbers = append(numbers, [2]int{ncur, nnex})
						}
					}
				}
			} else {
				// odd
				o := get_op(bcur)
				if o > 0 {
					l := numbers[k]
					y := k / 2
					if y >= 0 {
						numbers[y][k%2] = o.Calculate(l[0], l[1])
					}
				}

				// last operator
				// remove all redunant numbers slice
				if k+1 >= len(v) {
					numbers = numbers[:k+1]
				}
			}
		}
	}

	return numbers[0][0]
}

func main() {
	inputs := []string{
		`    *
   / \
  +    +
 / \  / \
3  2  4  5`,
	}

	for _, v := range inputs {
		fmt.Println(v, "==", solve(v))
	}
}

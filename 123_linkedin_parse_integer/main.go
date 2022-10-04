package main

import (
	"fmt"
	"math"
	"strings"
)

var numMap = map[byte]int{
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

func parse(str string) *float64 {
	if len(str) == 0 {
		// empty string
		return nil
	}

	negative := str[0] == '-'
	if negative {
		if len(str) == 1 {
			// minus without number
			return nil
		} else {
			str = str[1:]
		}
	}

	val := float64(0)
	decimalPoint := -1
	bytes := []byte(str)
	for i := range str {
		if str[i] == '.' {
			decimalPoint = i
		} else if str[i] == 'e' {
			if len(str) > i+1 {
				numChar := str[i+1]
				num, ok := numMap[numChar]
				if !ok {
					// e but character afterwards isn't number
					return nil
				}

				bytes = append(bytes[:i], append([]byte(strings.Repeat("0", num)), bytes[i+2:]...)...)
			} else {
				// e without number afterwards
				return nil
			}
		}
	}
	str = string(bytes)

	for i := range str {
		var zeros float64
		if i >= decimalPoint && decimalPoint != -1 {
			zeros = 10 / math.Pow(10, float64((i+1)-decimalPoint))
		} else {
			max := len(str)
			// if number is real
			if decimalPoint != -1 {
				max = decimalPoint
			}

			zeros = math.Pow(10, float64(max-(i+1)))
		}

		_, ok := numMap[str[i]]
		if decimalPoint != i && !ok {
			// not a number and not a decimal point either
			return nil
		}
		val += float64(numMap[str[i]]) * zeros
	}

	if negative {
		val = val * -1
	}

	return &val
}

func main() {
	str := []string{
		"10",
		"-10",
		"10.1",
		"-10.1",
		"1e5",
		"a",
		"x 1",
		"a -2",
		"-",
	}

	nums := []float64{
		10,
		-10,
		10.1,
		-10.1,
		1e5,
	}

	fmt.Println(parse("10.1"))
	for i, v := range str {
		var want *float64
		var have *float64
		if len(nums) > i {
			want = &nums[i]
		}

		have = parse(v)
		strs := []string{v, "<nil>", "<nil>"}
		if want != nil {
			strs[2] = fmt.Sprintf("%f", *want)
		}
		if have != nil {
			strs[1] = fmt.Sprintf("%f", *have)
		}

		fmt.Println(strings.Join(strs, " "))
	}
}

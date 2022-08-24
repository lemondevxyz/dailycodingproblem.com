package main

import (
	"fmt"
	"strconv"
)

var digits = map[int][]string{
	2: {"a", "b", "c"},
	3: {"d", "e", "f"},
	4: {"g", "h", "i"},
	5: {"j", "k", "l"},
	6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"},
	8: {"t", "u", "v"},
	9: {"w", "x", "y", "z"},
}

func get_pattern(one, two []string) []string {
	//ret := []string{}
	if len(one) == 0 {
		return two
	} else if len(two) == 0 {
		return one
	} else if len(two) == 0 && len(one) == len(two) {
		return nil
	}

	ret := []string{}
	for _, s := range one {
		for _, v := range two {
			ret = append(ret, s+v)
		}
	}

	return ret
}

// so i was pretty much lost on how to do this
// but i finally came to this beautiful simple result.
// which is replace the array in-place
// example:
// arr:
// 	0: a, b, c
// 	1: d, e, f
//  2: g, h, i
// after mod:
// arr:
// 	0: a, b, c
// 	1: ad, ae, af, bd, be, bf, cd, ce, cf
// 	2: adg adh adi aeg aeh aei afg afh afi bdg bdh bdi beg beh bei bfg bfh bfi cdg cdh cdi ceg ceh cei cfg cfh cfi
// hopefully this is optimized, but i don't think so :9
func solve(num int) []string {
	arr := [][]string{}
	s := strconv.Itoa(num)
	for _, v := range s {
		n, _ := strconv.Atoi(string(v))
		arr = append(arr, digits[n])
	}

	//pt := []string{}
	for k := range arr {
		if k > 0 {
			one, two := arr[k-1], arr[k]
			arr[k] = get_pattern(one, two)
		}
	}

	if len(arr) == 0 {
		return nil
	}

	return arr[len(arr)-1]
}

func main() {
	ex := []int{23, 5, 9, 75}
	for _, v := range ex {
		fmt.Println(v, solve(v))
	}
}

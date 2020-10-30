package main

var cases_with_order = map[string]bool{
	"()":     true,
	"))":     false,
	"((":     false,
	"(())":   true,
	"()) (":  false,
	"() ()":  true,
	"(char)": true,
	")char(": false,
}

func solve_with_order(str string) bool {
	lindex := []int{}
	// key is lindex, value is value
	var rcount int
	rmap := map[int]int{}

	for i, v := range str {
		if v == '(' {

			lindex = append(lindex, i)
			r := find_close_parenthesis(str[i+1:])

			if r >= 0 {
				rmap[i] = r + i
			}

		} else if v == ')' {
			rcount++
		}
	}

	// not enough parenthesis
	if len(lindex) != rcount {
		return false
	}

	for _, v := range lindex {
		_, ok := rmap[v]
		if !ok {
			return false
		}
	}

	return true
}

func find_close_parenthesis(str string) int {
	var count int
	for i, v := range str {
		if v == '(' {
			count++
			continue
		} else if v == ')' {
			if count > 0 {
				count--
				continue
			} else if count == 0 {
				return i
			}
		}
	}

	return -1
}

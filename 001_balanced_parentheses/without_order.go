package main

// map[function_parameter]wanted_result
var cases_without_order = map[string]bool{
	"()":   true,
	"(()":  false,
	"())":  false,
	"(())": true,
	"((((": false,
	")(()": true,
	// invalid case
	// "(()))": true,
}

func solve_without_order(str string) bool {

	//var lindex, rindex = []int{}, []int{}
	var lcount, rcount int

	// for k, v := range str {
	for _, v := range str {
		if v == '(' {
			//lindex = append(lindex, k)
			lcount++
		} else if v == ')' {
			//rindex = append(rindex, k)
			rcount++
		}
	}

	//return len(lindex) == len(rindex)

	return lcount == rcount
}

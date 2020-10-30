package main

import (
	"fmt"
	"os"
)

//	    map[have]want{}
var cases = map[string]bool{
	"()":   true,
	"(()":  false,
	"())":  false,
	"(())": true,
	"((((": false,
	")(()": true,
	// invalid case
	// "(()))": true,
}

func main() {
	for k, v := range cases {
		if solve(k) != v {
			fmt.Printf("solve(\"%s\") != %t\n", k, v)
			os.Exit(1)
		}
	}
}

func solve(str string) bool {

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

package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("testing cases with order")
	for k, v := range cases_with_order {
		if solve_with_order(k) != v {
			fmt.Printf("solve(\"%s\") != %t\n", k, v)
			os.Exit(1)
		}
	}
	fmt.Println("success")

	fmt.Println("testing cases without order")
	for k, v := range cases_without_order {
		if solve_without_order(k) != v {
			fmt.Printf("solve(\"%s\") != %t\n", k, v)
			os.Exit(1)
		}
	}
	fmt.Println("success")

}

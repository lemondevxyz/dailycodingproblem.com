package main

import "fmt"

func solve(a []int, k int) {
	// amount of operations := len(a)-k
	for i := 0; i < len(a)-k; i++ {
		//fmt.Println(a[i], a[k+i])
		a[i], a[k+i] = a[k+i], a[i]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr)
	solve(arr, 2)
	fmt.Println(arr)
}

// i really hate this problem
package main

import "fmt"

func solve(arr []int, k int) []int {

	// invalid parameter
	if k > len(arr) || k <= 1 {
		return arr
	}

	// which arr value to compare to
	bef := 0
	aft := k

	for i := 0; i < len(arr); i++ {
		if i >= k {
			bef++
		}

		if i >= k && aft < len(arr)-1 {
			aft++
		}

		v := arr[i]
		if arr[bef] < v {
			arr[bef] = v
		}

		// index out of range
		if arr[aft] < v && aft != k {
			arr[aft] = v
		}

	}

	return arr[:(len(arr)-k)+1]
}

func main() {

	fmt.Println(solve([]int{10, 5, 2, 7, 8, 7}, 3))
	fmt.Println(solve([]int{5, 2, 1}, 2))

}

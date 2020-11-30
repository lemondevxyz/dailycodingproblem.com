package main

import "fmt"

func solve(arr []int) int {

	// naive attempt lamo
	/*
		x := 0
		for i := 0; i < len(arr); i++ {
			bef := i - 1
			if bef == -1 {
				bef = len(arr) - 1
			}

			aft := i + 1
			if aft >= len(arr) {
				aft = 0
			}

			if x == 0 {
				if arr[aft] != arr[i] && arr[bef] != arr[i] {
					x = arr[i]
				}
			} else if x == arr[i] {
				if arr[aft] == arr[i] || arr[bef] == arr[i] {
					x = 0
				} else {
					x = arr[i]
				}
			}
		}
	*/
	/*
			x := 0
				for i := 0; i < len(arr); i++ {
					bef := i - 1
					if bef == -1 {
						bef = len(arr) - 1
					}

					aft := i + 1
					if aft >= len(arr) {
						aft = 0
					}

					if x == arr[i] {
						continue
					}
					// don't swap if it's 3 in a row
					if arr[i] == arr[bef] && arr[i] != arr[aft] {
						ol := arr[aft]
						arr[aft] = arr[bef]
						arr[bef] = ol
						x = arr[i]
					} else if arr[i] == arr[aft] && arr[i] != arr[bef] {
						ol := arr[bef]
						arr[bef] = arr[aft]
						arr[aft] = ol
						x = arr[i]
					}
				}
		/*
			for i := 0; i < 3; i++ {
				v := i * 3
				fmt.Println(arr[v : v+3])
			}*/
	for i := 0; i < len(arr); i++ {
		befn := i - 1
		if befn == -1 {
			befn = len(arr) - 1
		}
		aftn := i + 1
		if aftn >= len(arr) {
			aftn = 0
		}

		aft := arr[aftn]
		bef := arr[befn]
		cur := arr[i]

		if aft != cur && bef != cur {
			arr[aftn] = cur
			arr[i] = aft
		}
	}

	return arr[0]
}

func main() {
	arrs := [][]int{
		{6, 1, 3, 3, 3, 6, 6},
		{13, 19, 13, 13},
		{9, 9, 9, 2, 5, 5, 5},
		{5, 2, 6, 9, 5, 5, 6, 6, 9, 9},
	}
	for _, v := range arrs[:4] {
		arr := []int{}
		arr = append(arr, v...)
		//arr = v
		fmt.Printf("%v: %d\n", arr, solve(v))
	}
}

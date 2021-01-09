package main

import "fmt"

func solve(matrix [][]rune, word string) bool {
	if len(matrix) == 0 || len(matrix) != len(word) {
		return false
	}

	sli := []string{}
	for i := 0; i < len(matrix); i++ {
		sli = append(sli, string(matrix[i]))

		arr := []rune{}
		for j := 0; j < len(matrix[i]); j++ {
			arr = append(arr, matrix[j][i])
		}
		sli = append(sli, string(arr))
	}

	for _, v := range sli {
		if v == word {
			return true
		}
	}

	return false
}

func main() {
	m := [][]rune{
		[]rune{'F', 'A', 'S', 'I'},
		[]rune{'O', 'B', 'A', 'P'},
		[]rune{'A', 'N', 'A', 'B'},
		[]rune{'M', 'A', 'S', 'S'},
	}

	for _, v := range m {
		for _, s := range v {
			fmt.Printf("%c ", s)
		}
		fmt.Println()
	}
	//fmt.Println(m)

	words := []string{"FOAM", "MASS", "SAAS", "BLUE"}
	for _, v := range words {
		fmt.Printf("%s: %t\n", v, solve(m, v))
	}

}

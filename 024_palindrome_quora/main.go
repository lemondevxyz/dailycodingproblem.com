package main

import "fmt"

func reverse(str string) string {
	rev := ""

	for _, v := range str {
		rev = string(v) + rev
	}

	return rev
}

func isPalindrome(str string) bool {
	return str == reverse(str)
}

// O(n^2) time
func solve(str string) string {
	if len(str) == 0 {
		return ""
	}

	// dad
	if isPalindrome(str) {
		return str
	}

	// first reverse the string
	rev := reverse(str)

	for i := 0; i < len(rev); i++ {
		result := rev[:i+1] + str
		if isPalindrome(result) {
			return result
		}
	}

	return ""
}

func main() {
	words := []string{"race", "google", "dad"}

	for _, v := range words {
		fmt.Println(v+":", solve(v))
	}
}

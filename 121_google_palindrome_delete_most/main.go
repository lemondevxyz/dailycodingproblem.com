package main

import "fmt"

func reverse(s string) string {
	arr := []byte(s)
	ret := []byte{}

	for i := len(arr) - 1; i >= 0; i-- {
		ret = append(ret, arr[i])
	}

	return string(ret)
}

func isPalindrome(str string) bool {
	isOdd := (len(str) % 2) == 1
	size := len(str) / 2

	if isOdd {
		return str[0:size] == reverse(str[size+1:])
	}
	return str[0:size] == reverse(str[size:])
}

func subslice(arr []byte, i int) []byte {
	if i == 0 {
		return arr[1:]
	} else if i+1 > len(arr) {
		return arr[:len(arr)-1]
	}

	return append(arr[0:i], arr[i+1:]...)
}

// this works through recursion
// basically, solve looks for each character in word
// it removes one character, and checks if its a palindrome
// if it is, it returns the word
// otherwise, it doesn't, it goes through the other characters
func solve(word string, k int) string {
	for x := 0; x < len(word); x++ {
		for i := 0; i < k; i++ {
			myword := string(subslice([]byte(word), x))
			if isPalindrome(myword) {
				return myword
			}
			solveWord := solve(myword, k-1)
			if len(solveWord) > 0 {
				return solveWord
			}
		}
	}

	return ""
}

func main() {
	fmt.Println("waterfetawx", solve("waterrfetawx", 2))
}

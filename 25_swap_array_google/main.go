package main

import (
	"bytes"
	"fmt"
)

func solve(chars []byte) []byte {
	count := [3]int{}
	for _, v := range chars {
		if v == 'R' {
			count[0]++
		} else if v == 'G' {
			count[1]++
		} else if v == 'B' {
			count[2]++
		}
	}

	result := []byte{}
	result = append(result, bytes.Repeat([]byte{'R'}, count[0])...)
	result = append(result, bytes.Repeat([]byte{'G'}, count[1])...)
	result = append(result, bytes.Repeat([]byte{'B'}, count[2])...)

	return result
}

func main() {
	input := [][]byte{
		{'G', 'B', 'R', 'R', 'B', 'R', 'G'},
	}

	for _, v := range input {
		fmt.Println(string(v)+":", string(solve(v)))
	}
}

package main

import (
	"fmt"
	"os"
)

type Matrix [][]bool

func sum(arr []bool) (i int) {
	for _, v := range arr {
		if v {
			i++
		}
	}

	return
}

// meaning input will be a matrix and k, output will be a boolean value.
// well then just calculate max sum of row, then comparei t to k
func solve(m Matrix, k int) bool {
	max := 0
	for _, v := range m {
		x := sum(v)
		if max < x {
			max = x
		}
	}

	return k > max
}

type Case struct {
	have bool
	want bool
}

var cases = []Case{

	// matrix 1

	{
		have: solve(Matrix{
			{false, true, true, true},
			{true, false, true, true},
			{true, true, false, true},
			{true, true, true, false},
		}, 4),
		want: true,
	},
	{
		have: solve(Matrix{
			{false, true, true, true},
			{true, false, true, true},
			{true, true, false, true},
			{true, true, true, false},
		}, 3),
		want: false,
	},

	// matrix 2

	{
		have: solve(Matrix{
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		}, 4),
		want: true,
	},

	{
		have: solve(Matrix{
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		}, 1),
		want: true,
	},
	// matrix 3
	{
		have: solve(Matrix{
			{false, true, true, false},
			{true, false, false, true},
			{true, false, false, true},
			{false, true, true, false},
		}, 4),
		want: true,
	},
	{
		have: solve(Matrix{
			{false, true, true, false},
			{true, false, false, true},
			{true, false, false, true},
			{false, true, true, false},
		}, 3),
		want: true,
	},
	{
		have: solve(Matrix{
			{false, true, true, false},
			{true, false, false, true},
			{true, false, false, true},
			{false, true, true, false},
		}, 2),
		want: false,
	},
}

func main() {
	for k, v := range cases {
		if v.have != v.want {
			fmt.Printf("bad code. index %d\n", k)
			os.Exit(1)
		}
	}

	fmt.Println("success")
}

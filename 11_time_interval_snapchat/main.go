// no comments edition
package main

import (
	"fmt"
	"os"
	"sort"
)

type Interval struct {
	start int
	end   int
}

type AllInterval []Interval

func (a AllInterval) Len() int { return len(a) }

func (a AllInterval) Less(i, j int) bool {
	return a[i].start < a[j].start
}

func (a AllInterval) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func solve(arr AllInterval) (result int) {

	if len(arr) == 0 {
		return 0
	}

	sort.Sort(arr)

	for k := range arr {
		l := k + 1
		if l == len(arr) {
			l = 0
		}

		v := arr[k]
		w := arr[l]

		if v.start <= w.end {
			result++
		}
	}

	if result == 0 && len(arr) > 0 {
		result = 1
	}

	return result

}

type Case struct {
	want int
	have int
}

var cases = []Case{
	{
		have: solve([]Interval{
			{
				start: 30,
				end:   75,
			},
			{
				start: 0,
				end:   50,
			},
			{
				start: 60,
				end:   150,
			},
		}),
		want: 2,
	},

	{
		have: solve([]Interval{
			{
				start: 30,
				end:   75,
			},
			{
				start: 0,
				end:   50,
			},
			{
				start: 60,
				end:   150,
			},
			{
				start: 10,
				end:   60,
			},
		}),
		want: 3,
	},

	{
		have: solve([]Interval{
			{
				start: 60,
				end:   150,
			},
		}),
		want: 1,
	},

	{
		have: solve([]Interval{
			{
				start: 60,
				end:   150,
			},
			{
				start: 150,
				end:   170,
			},
		}),
		want: 2,
	},

	{
		have: solve([]Interval{
			{
				start: 60,
				end:   150,
			},
			{
				start: 60,
				end:   150,
			},
			{
				start: 150,
				end:   170,
			},
		}),
		want: 3,
	},
}

func main() {

	fmt.Println(cases)
	for k, v := range cases {
		if v.want != v.have {
			fmt.Printf("want: %d, have: %d\n", v.want, v.have)
			fmt.Println("index", k)
			os.Exit(1)
		}
	}
}

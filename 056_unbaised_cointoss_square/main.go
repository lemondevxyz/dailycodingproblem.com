package main

import (
	"fmt"
	"math/rand"
	"time"
)

var prob = int(time.Now().UnixNano())

func solve() int {
	gt := rand.Intn(prob)
	num := rand.Intn(prob)

	if num > gt {
		return 1
	}

	return 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	total := 1000
	mx := map[int]int{}

	for i := 0; i < total; i++ {
		mx[solve()]++
	}
	perc1 := float64(mx[0]*100) / float64(total)
	perc2 := float64(mx[1]*100) / float64(total)

	fmt.Println(perc1, perc2)
}

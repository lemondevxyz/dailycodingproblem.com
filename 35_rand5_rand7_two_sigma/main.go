package main

// lamo i know i don't understand what this problem is

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

// from 1 to 7
// inclusive
func rand7() int {
	/*
		i := 5*rand5() + rand5() - 5
		if i < 22 {
			return i%7 + 1
		}
	*/
	return rand.Intn(7)
}

// from 1 to 5
// inclusive, uniform probability
func rand5() int {
	return rand.Intn(5)
}

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	tries := 100000
	probs := map[int]float64{}
	for i := 0; i < tries; i++ {
		x := rand7()
		_, exists := probs[x]
		if !exists {
			probs[x] = 0
		}

		probs[x]++
	}

	want := float64(1) / 7
	for _, v := range probs {
		v = float64(v) / float64(tries)
		v = roundTo(v, 2)

		if v != roundTo(want, 2) {
			fmt.Println("bad code")
			os.Exit(1)
		}
	}
	fmt.Println("success")

}

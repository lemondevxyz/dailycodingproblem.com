package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

// i never understand these
// lamo just use the default library
func solve(k int) int {
	return rand.Intn(k)
}

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func main() {

	roundAmount := uint32(2)
	k := 52
	tries := 10000

	possib := float64(1) / 52
	possib = roundTo(possib, roundAmount)

	rand.Seed(time.Now().UTC().UnixNano())
	gene := map[int]int{}

	for i := 0; i < tries; i++ {
		res := solve(k)
		_, ok := gene[res]
		if !ok {
			gene[res] = 0
		}

		gene[res]++
	}

	for _, v := range gene {
		perc := float64(v) / float64(tries)
		perc = roundTo(perc, roundAmount)
		if perc != possib {
			fmt.Println("bad random")
			os.Exit(1)
		}
	}
	fmt.Println("success")
}

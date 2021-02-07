package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func solve(n int, l []int) int {
	m := map[int]struct{}{}
	for _, v := range l {
		m[v] = struct{}{}
	}

	y := []int{}
	for i := 0; i < n; i++ {
		_, ok := m[i]
		if !ok {
			y = append(y, i)
		}
	}

	x := rand.Intn(len(y))

	return y[x]
}

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func main() {

	rand.Seed(time.Now().UnixNano())

	tries := 1000

	n := 10
	l := []int{5, 8, 10}

	stuff := map[int]int{}

	for i := 0; i < tries; i++ {
		x := solve(n, l)
		stuff[x]++
	}

	for _, v := range l {
		_, ok := stuff[v]
		if ok {
			log.Fatalf("bad code")
		}
	}
	fmt.Println("ehh success")

}

// lamo this one was hard. i understood what monte carlo was about but didn't know why it mattered.
// later on watched a video and everything clicked.
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// n = amount of random tries
func solve(n int) float64 {

	getXY := func() (float64, float64) {
		return rand.Float64(), rand.Float64()
	}

	getR := func() float64 {
		x, y := getXY()

		// x^2 + y^2 = r
		return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	}

	inside := 0

	for i := 0; i < n; i++ {

		r := getR()
		if r <= 1 {
			inside++
		}

	}

	// yay static types
	div := float64(inside) / float64(n)

	return div * 4
}

func main() {
	fmt.Println(solve(1000))
}

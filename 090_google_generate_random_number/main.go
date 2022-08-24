package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func solve(n int, l []int) (v int) {
	exists := true
	for exists {
		v = rand.Intn(n)
		exists = false
		for _, x := range l {
			if x == v {
				exists = true
			}
		}
	}

	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		x := solve(14, []int{2, 5, 7})
		if x == 2 || x == 5 || x == 7 {
			log.Fatalf("henlo", x)
		}
	}

	fmt.Println("werks")
}

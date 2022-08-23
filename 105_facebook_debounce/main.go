package main

import (
	"fmt"
	"time"
)

func debounce(f func(), n time.Duration) func() {
	start := time.Now()
	return func() {
		now := time.Now()

		sleep := now.Sub(start)
		if sleep > n {
			sleep = 0
		} else {
			sleep = n - sleep
		}
		time.Sleep(sleep)

		f()
		start = time.Now()
	}
}

func main() {
	for i := 0; i < 5; i++ {
		j := i

		fn := debounce(func() {
			fmt.Println(j)
		}, time.Second)
		fn()

	}
}

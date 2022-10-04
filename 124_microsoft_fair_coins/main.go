package main

// note i have no idea how this works
// seriously, wtf

// no way we can predict and any guess is as good as any
import "fmt"

func solve(n int) int {
	tails := n
	rounds := 0

	for tails > 1 {
		tails = tails / 2
		rounds++
	}

	return rounds
}

func main() {
	fmt.Println("Coins:", 15)
	fmt.Println("Amount of rounds before coins:", solve(15))
}

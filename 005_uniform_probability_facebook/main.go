package main

// at first i misunderstood what was needed. i thought i would have an array and need to pick a random value inside that array. simple enough just do math.Intn(len(arr))
// then after reading multiple articles on this issue I understood it very well. basically you're required to return a random element from an array. but you don't have said array. you only have one of the elements and it wants the randomness to be of uniform probability.
// meaning it's a 1 in N chance that you pick a random element.
// you could do this two ways, one by counting the amount of elements you are currently having and then generating a random float. then multiply that float by count. call this variable `val`
// now compare val either to the biggest number or smallest number, so compare it to `count` or to `0`
// if it's the comparison is valid then set the global variable to your current element
// the second way you could do this is comparing the actual floats, set a global variable named `lastprob`, then in the function set a local variable `prob`. copare probs to be greater or less than lastprob(if you want to compare it to be less you have to set lastprob initially to 1). if the condition succeeds return current element, else return the last random element
// thanks to fellow developer Nikita Galaiko(https://galaiko.rocks) for his amazing explaination on how to solve this problem. bonus: he solves problems in Go too :).

import (
	"fmt"
	"math/rand"
	"time"
)

var lastprob float64
var result int

func solve(element int) int {
	rand.Seed(time.Now().UTC().UnixNano())

	prob := rand.Float64()

	if prob > lastprob {
		lastprob = prob
		result = element
	}

	return result
}

func main() {

	for i := 0; i < 9; i++ {
		fmt.Println(solve(i))
	}
}

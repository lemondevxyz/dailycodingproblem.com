package main

import "fmt"

// this algorithm assumes that the length represents the order of the courses
// furthermore, there is no cross matching. therefore this algorithm is dumb
// if provided with more examples, I could've extended this algorithm to support cross matching
func solve(m map[string][]string) []string {
	sl := map[int]string{}
	for k, v := range m {
		sl[len(v)] = k
	}

	ret := []string{}
	for i := 0; i < len(sl); i++ {
		ret = append(ret, sl[i])
	}

	return ret
}

func main() {
	fmt.Println(solve(map[string][]string{
		"CSC300": []string{"CSC200", "CSC100"},
		"CSC200": []string{"CSC100"},
		"CSC100": nil,
	}))
}

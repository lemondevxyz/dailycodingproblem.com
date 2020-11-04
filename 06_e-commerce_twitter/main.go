package main

import "fmt"

func main() {
	samples := []string{"ab", "ac", "ad", "ba", "bc", "bd", "ca", "cb", "cd", "da", "db", "dc"}

	for _, v := range samples {
		record(v)
	}

	if get_last(2).id == "ad" &&
		get_last(4).id == "bc" &&
		get_last(6).id == "ca" {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}

var orders = []order{}

type order struct {
	id string
}

func record(id string) int {

	index := len(orders)
	orders = append(orders, order{id: id})

	return index
}

func get_last(i int) order {
	return orders[i]
}

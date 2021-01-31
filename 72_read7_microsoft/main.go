package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

func read7(file io.Reader) ([]byte, error) {
	b := make([]byte, 7)

	n, err := file.Read(b)
	if err != nil {
		return nil, err
	}

	return b[:n], nil
}

func readN(file io.Reader, n int) ([]byte, error) {

	x := n / 7
	fl := float64(n) / 7
	if math.Mod(fl, 1) != 0 {
		x++
	}

	b := []byte{}
	for i := 0; i < x; i++ {
		bod, err := read7(file)
		if err != nil {
			return nil, err
		}

		b = append(b, bod...)
	}

	return b[:n], nil
}

func main() {
	str := "Hello World"

	for i := 1; i <= len(str); i++ {
		r := strings.NewReader(str)
		bod, err := readN(r, i)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(bod))
	}
}

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
)

func solve(want, have []byte, size int) {
	pos := 0
	wr, hr := bytes.NewReader(want), bytes.NewReader(have)
	for {
		w, h := make([]byte, size), make([]byte, size)
		wn, we := wr.Read(w)
		_, he := hr.Read(h)
		if we == io.EOF && he == nil {
			// remove data from have
			have = have[:wn]
		} else if we == nil && he == io.EOF {
			// add data to have
			have = append(have, w...)
		} else if we == io.EOF && he == io.EOF {
			break
		}

		ws, hs := sha256.Sum256(w), sha256.Sum256(h)
		if !bytes.Equal(ws[:], hs[:]) {
			end := pos + size
			if end > len(want) {
				end = len(want)
			}

			copy(have[pos:end], want[pos:end])
		}
		pos += size
	}
}

func main() {
	byt1 := []byte("ab cd ef")
	byt2 := []byte("ab cd gh")

	fmt.Printf("%s\n%s\n", string(byt1), string(byt2))
	solve(byt1, byt2, 1024)
	fmt.Printf("%s\n%s\n", string(byt1), string(byt2))
}

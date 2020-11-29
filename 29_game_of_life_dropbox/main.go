package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func generateRandomBoard(size int) [][]bool {
	b := [][]bool{}
	for i := 0; i < size; i++ {
		b = append(b, []bool{})
		for j := 0; j < size; j++ {
			bvalue := (rand.Intn(2) == 1)

			b[len(b)-1] = append(b[len(b)-1], bvalue)
		}
	}

	return b
}

type board struct {
	data [][]bool
}

type cell struct {
	i int
	j int
}

func newBoard(size int) board {

	b := board{
		data: generateRandomBoard(size),
	}

	return b

}

func (b board) String() string {
	str := ""
	for _, v := range b.data {
		for _, c := range v {
			char := "."
			if c {
				char = "*"
			}
			str += char + " "
		}
		str += "\n"
	}

	return str
}

func safebef(i, n int) int {
	i = i - 1
	if i < 0 {
		i = n - 1
	}

	return i
}

func safeaft(i, n int) int {
	i++
	if i >= n {
		i = 0
	}

	return i
}

// neighbors are O in following
// NW N NE
//  W C E
// SW S SE
// live cell with less than 2 live neighbours dies

func (b board) neighbours(i, j int) []cell {
	size := len(b.data)
	befi := safebef(i, size)
	afti := safeaft(i, size)
	befj := safebef(j, size)
	aftj := safebef(j, size)

	nw := cell{befj, befi}
	n := cell{befi, j}
	ne := cell{befi, aftj}

	w := cell{i, befj}
	e := cell{i, aftj}

	sw := cell{befj, afti}
	s := cell{afti, j}
	se := cell{afti, aftj}

	return append([]cell{}, nw, n, ne, w, e, sw, s, se)
}

func (b board) Tick() {

	n := len(b.data)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cells := b.neighbours(i, j)
			alive := 0
			for _, v := range cells {
				if b.data[v.i][v.j] {
					alive++
				}
			}

			thiscell := b.data[i][j]
			if thiscell {
				if alive < 2 || alive > 3 {
					b.data[i][j] = false
				} else {
					b.data[i][j] = true
				}
			} else {
				if alive == 3 {
					b.data[i][j] = true
				}
			}
		}
	}

}

// live update
const ESC = 27

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	size := 7

	b := newBoard(size)
	clear := fmt.Sprintf("%c[%dA%c[2K", ESC, 1, ESC)
	fmt.Println("Press CTRL-C to stop")
	fmt.Print(b)
	// for men

	go func() {
		for {
			time.Sleep(time.Second * 1)
			b.Tick()
			fmt.Print(strings.Repeat(clear, size))
			fmt.Print(b)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}

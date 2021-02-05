package main

import (
	"fmt"
	"strings"
)

type dir uint8

const (
	North dir = iota + 1
	South
	East
	West
	NorthEast
	NorthWest
	SouthEast
	SouthWest
)

var stringdir = map[dir]string{
	North:     "N",
	South:     "S",
	East:      "E",
	West:      "W",
	NorthEast: "NE",
	NorthWest: "NW",
	SouthEast: "SE",
	SouthWest: "SW",
}

type Rule struct {
	src string
	dir dir
	dst string
}

func (r Rule) Goes(d dir) bool {
	oppo := map[dir][]dir{
		NorthEast: {North, East},
		NorthWest: {North, West},
		SouthWest: {South, West},
		SouthEast: {South, East},
	}

	for _, v := range oppo[r.dir] {
		if v == d {
			return true
		}
	}

	return false
}

func (r Rule) Opposite() Rule {
	oppo := map[dir]dir{
		North:     South,
		South:     North,
		East:      West,
		West:      East,
		NorthEast: SouthWest,
		NorthWest: SouthEast,
		SouthWest: NorthEast,
		SouthEast: NorthWest,
	}

	o := r
	o.src, o.dst = r.dst, r.src

	o.dir = oppo[r.dir]

	return o
}

func (r Rule) String() string {
	return fmt.Sprintf("%s %s %s", r.src, stringdir[r.dir], r.dst)
}

func ParseRule(str string) Rule {

	spl := strings.Split(str, " ")

	r := Rule{}
	for k, v := range spl {
		if k == 0 {
			// first
			r.src = v
		} else if k+1 == len(spl) {
			// last
			r.dst = v
		} else {
			// middle
			for k, d := range stringdir {
				if v == d {
					r.dir = k
					break
				}
			}
		}
	}

	return r
}

func solve(str []string) bool {
	rs := []Rule{}
	for _, v := range str {
		r := ParseRule(v)

		rs = append(rs, r)
	}

	for _, v := range rs {
		for _, s := range rs {
			if v.src == s.dst && v.dst == s.src {
				if v.Opposite().dir != s.dir {
					if !v.Goes(s.dir) {
						return false
					}
				}
			}
		}
	}

	return true
}

func main() {
	s := []string{
		"A N B",
		"A NE C",
		"C N A",
	}

	fmt.Println(strings.Join(s, "\n"), solve(s))

	s = []string{
		"A NW B",
		"A N B",
	}

	fmt.Println(strings.Join(s, "\n"), solve(s))
}

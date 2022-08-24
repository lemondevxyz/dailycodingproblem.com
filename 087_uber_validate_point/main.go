package main

import (
	"fmt"
	"strings"
	"time"
)

// Strategy: loop through every rule
// then loop through every other rule in the midst of the first loop
// generate a list of invalid rules
// compare the original rule to the invalid rule
// if it exists, return false
type rule struct {
	src byte
	dst byte
	dir direction
}

func (r rule) Equal(o rule) bool {
	return r.src == o.src && r.dst == o.dst && r.dir == o.dir
}

func (r rule) String() string {
	return fmt.Sprintf("%c %s %c", r.src, directionmap[r.dir], r.dst)
}

// return a list of rules with every other direction
func (r rule) Dir() []rule {
	s := []rule{}
	for _, v := range dirs {
		if v == r.dir || mergable(v, r.dir) > 0 {
			continue
		}

		s = append(s, rule{
			r.src,
			r.dst,
			v,
		})
	}

	return s
}

func (r rule) Inverse() rule {
	return rule{r.dst, r.src, inversemap[r.dir]}
}

func (r rule) Possib() []rule {
	return append(r.Dir(), r.Inverse().Dir()...)
}

type direction uint8

const (
	north direction = 1 << iota
	south
	west
	east

	northwest = north | west
	northeast = north | east
	southwest = south | west
	southeast = south | east
)

var inversemap = map[direction]direction{north: south, south: north, east: west, west: east, northwest: southeast, northeast: southwest, southeast: northwest, southwest: northeast}
var directionmap = map[direction]string{}
var stringmap = map[string]direction{"N": north, "S": south, "W": west, "E": east, "NW": northwest, "NE": northeast, "SW": southwest, "SE": southeast}
var dirs = []direction{north, south, west, east, northwest, northeast, southwest, southeast}

func init() {
	for k, v := range stringmap {
		directionmap[v] = k
	}
}

func parse(str string) rule {
	spl := strings.Split(str, " ")
	return rule{spl[0][0], spl[2][0], stringmap[spl[1]]}
}

func exists(r rule, slice []rule) bool {
	for _, v := range slice {
		if v.Equal(r) {
			return true
		}
	}

	return false
}

func mergable(v, o direction) direction {
	if v == o {
		return v
	}

	// i.e NorthEast North
	// or NorthEast East
	if (v | o) != 0 {
		val := direction(0)
		if v&^o > 0 {
			val = (v &^ o) ^ v
		} else if o&^v > 0 {
			val = (o &^ v) ^ o
		}

		return val
	}

	return 0
}

// Strategy: consider the following rules
// A N B
// B N C
// C N D
//
// Loop through every rule, then loop through every other rule.
// Try to find a rule that has src equal to current-rule's dst.
// If it does, add it, and restart.
// Until len(r) == len(ret)
func lead(r []rule, exist map[string]bool) (ret []rule) {
	for _, v1 := range r {
		for _, v2 := range r {
			if v1.dst == v2.src && mergable(v1.dir, v2.dir) > 0 {
				rl := rule{
					src: v1.src,
					dst: v2.dst,
					dir: v1.dir,
				}

				_, ok := exist[rl.String()]
				if ok {
					break
				}

				ret = append(ret, r...)
				ret = append(ret, rl)
				exist[rl.String()] = true

				return
			}
		}
	}

	return r
}

func lead_loop(r []rule) []rule {
	old, new := 0, len(r)
	m := map[string]bool{}
	for old != new {
		old = new
		r = lead(r, m)
		new = len(r)
		time.Sleep(time.Millisecond * 50)
	}

	return r
}

func except(r []rule, i int) []rule {
	if i == len(r) {
		return r[:i]
	}
	return append(r[:i], r[i+1:]...)
}

func invalid(str string) error {
	rules := []rule{}
	for _, v := range strings.Split(str, "\n") {
		rules = append(rules, parse(v))
	}

	for k, v := range rules {
		rules2 := append([]rule{}, rules[:k]...)

		rules2 = lead_loop(rules2)
		rules3 := []rule{}
		for k := len(rules2) - 1; k >= 0; k-- {
			rules3 = append(rules3, rules2[k].Possib()...)
		}

		for _, r := range rules3 {
			if v.Equal(r) {
				return fmt.Errorf("error: %v is invalid", v)
			}
		}
	}

	return nil
}

func main() {
	isinvalid := invalid(`A N B
B NE C
C N A`)
	if isinvalid == nil {
		fmt.Println("want: <error>, have: nil")
	}

	isinvalid = invalid(`A NW B
A N B`)
	if isinvalid != nil {
		fmt.Println("want: nil, have:", isinvalid)
	}
}

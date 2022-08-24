package main

import "fmt"

type Edit struct {
	Type    uint
	Index   int
	OriChar byte
	NewChar byte
}

func (e Edit) String() string {
	if e.Type == 0 {
		// normal
		return fmt.Sprintf("substitute the %c for %c", e.OriChar, e.NewChar)
	} else if e.Type == 1 {
		// remove
		return fmt.Sprintf("remove %c", e.OriChar)
	} else if e.Type == 2 {
		// append
		return fmt.Sprintf("append %c", e.NewChar)
	}

	return ""
}

type Edits struct {
	arr    []Edit
	oristr string
	newstr string
}

func (es Edits) String() (str string) {

	for k, v := range es.arr {
		var char string

		if k+1 == len(es.arr) {
			char = "."
		} else {
			char = ", "
		}

		str += v.String() + char
	}

	return
}

func (es Edits) Apply() string {

	ori := []byte(es.oristr)
	for _, v := range es.arr {
		newchar := v.NewChar
		if v.Type == 1 {
			newchar = byte(0)
		}

		if v.Index < len(es.oristr) {
			ori[v.Index] = newchar
		} else {
			ori = append(ori, newchar)
		}
	}

	return string(ori)
}

// oristr is original
// newstr is new string
// meaning get different characters from oristr and replace em in newstr
func solve(oristr, newstr string) Edits {

	es := Edits{}
	es.oristr = oristr
	es.newstr = newstr

	for k := range newstr {
		if len(oristr) > k {
			if oristr[k] != newstr[k] {
				es.arr = append(es.arr, Edit{
					Type:    0,
					Index:   k,
					OriChar: newstr[k],
					NewChar: oristr[k],
				})
			}
		}
	}

	diff := len(newstr) - len(oristr)
	if diff > 0 {
		for i := len(newstr) - diff; i < len(newstr); i++ {
			es.arr = append(es.arr, Edit{
				// meaning remove
				Index:   i,
				Type:    1,
				OriChar: newstr[i],
			})
		}
	} else {
		diff = diff * -1
		for i := len(oristr) - diff; i < len(oristr); i++ {
			es.arr = append(es.arr, Edit{
				// meaning append
				Index:   i,
				Type:    2,
				NewChar: oristr[i],
			})
		}
	}

	return es
}

func main() {
	es := solve("kitten", "sitting")
	fmt.Println(es.Apply(), es)
	es = solve("sitting", "kitten")
	fmt.Println(es.Apply(), es)
}

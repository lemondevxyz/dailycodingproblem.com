package main

import (
	"fmt"
	"strings"
)

type Node struct {
	v int
	l *Node
	r *Node
}

type NodeItem struct {
	v int
	l bool
}

func (n NodeItem) Equal(o NodeItem) bool {
	return n.v == o.v && n.l == o.l
}

func (n NodeItem) String() string {
	return fmt.Sprintf("%d", n.v)
}

func (n *Node) do(arr [][]NodeItem, index int) ([][]NodeItem, int) {
	original := append([]NodeItem{}, arr[index]...)
	if n.l != nil {
		arr[index] = append(arr[index], NodeItem{n.l.v, false})
		arr, index = n.l.do(arr, index)
	}

	if n.l != nil && n.r != nil {
		arr = append(arr, []NodeItem{{n.v, false}})
		index++
	}

	if n.r != nil {
		arr[len(arr)-1] = append(append([]NodeItem{}, original...), NodeItem{n.r.v, true})
		arr, index = n.r.do(arr, index)
	}

	return arr, index

}

func (n *Node) string(i int) (str string) {
	str += fmt.Sprintf("%s%d\n", strings.Repeat(" ", i), n.v)

	if n.l != nil {
		str += "l" + n.l.string(i+1)
	}

	if n.r != nil {
		str += "r" + n.r.string(i+1)
	}

	return
}

func subarr(s []NodeItem) [][]NodeItem {
	ret := append([][]NodeItem{}, s)
	for i := 1; i < len(s); i++ {
		ret = append(ret, s[i:])
	}

	return ret
}

func equal(s1, s2 []NodeItem) bool {
	if len(s1) != len(s2) {
		return false
	}

	for k1, v1 := range s1 {
		if k1 == 0 {
			if s2[k1].v != v1.v {
				return false
			}
		} else {
			if !s2[k1].Equal(v1) {
				return false
			}
		}
	}

	return true
}

func solve(roots [][]NodeItem, sub [][]NodeItem) bool {
	oks := 0
	for _, v := range sub {
		tree := subarr(v)
		//eq := true
		for _, root := range roots {
			eq := false
			for _, nodeSlice := range tree {
				if equal(root, nodeSlice) {
					eq = true
				}
			}

			if eq {
				oks++
			}
		}

	}

	return oks == len(roots)
}

func main() {
	t := &Node{
		1,
		&Node{
			2,
			&Node{
				4,
				&Node{
					5, nil, nil,
				},
				nil,
			},
			&Node{
				3, nil, nil,
			},
		},
		&Node{
			6,
			nil,
			&Node{
				7,
				&Node{
					8,
					&Node{
						10, nil, nil,
					},
					&Node{
						9, nil, nil,
					},
				},
				nil,
			},
		},
	}

	s := &Node{
		0,
		&Node{
			1,
			&Node{
				2,
				&Node{
					4,
					&Node{
						5, nil, nil,
					},
					nil,
				},
				&Node{
					3, nil, nil,
				},
			},
			&Node{
				6,
				nil,
				&Node{
					7,
					&Node{
						8,
						&Node{
							10, nil, nil,
						},
						&Node{
							9, nil, nil,
						},
					},
					nil,
				},
			},
		},
		nil,
	}

	fmt.Println("first tree (t):")
	fmt.Println(t.string(0))

	fmt.Println("second tree (s):")
	fmt.Println(s.string(0))

	sNodes, _ := s.do([][]NodeItem{[]NodeItem{NodeItem{s.v, false}}}, 0)
	tNodes, _ := t.do([][]NodeItem{[]NodeItem{NodeItem{t.v, false}}}, 0)

	fmt.Println(solve(tNodes, sNodes))
}

// i liked this problem
package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type Node struct {
	Name     string
	IsFile   bool
	Children []*Node
}

const tabstr = "\t"
const tabrune = rune(9)

func count_tabs(str string) (amount int) {
	for _, v := range str {
		if v == tabrune {
			amount++
		}
	}

	return
}

// returns the string with the selected directory only and it's children, e.g.
// \n\tsubdir1\n\t\tfile.ext\n\tsubdir2\n\t\tfile2.ext
// given subdir1, it'll return
// \n\tsubdir1\n\t\tfile.ext
func get_substring(str, dir string) string {

	want := count_tabs(dir)

	// if the directory doesn't exist, invalid string
	index := strings.Index(str, dir)
	if index == -1 {
		return ""
	}

	// skip anything behind our directory
	str = str[index:]
	split := strings.Split(str, "\n")

	// last key in split
	index = 0

	if len(split) == 0 {
		return ""
	}

	// skip the parent
	for k, v := range split[1:] {
		have := count_tabs(v)
		if have > want {
			// cause we skipped the parent
			index = k + 1
		} else {
			// we've reached the next directory in level
			break
		}
	}

	// put it back together
	// reason we add +1, is because we want the value of index
	// if it was split[:index], we wouldn't have
	// split[index] as a value

	return strings.Join(split[:index+1], "\n")

}

func scan_string(str string) (n Node) {

	s := strings.Split(str, "\n")
	if len(s) == 0 {
		return
	}
	parent := s[0]

	n.Name = strings.TrimSpace(parent)
	//isourcase := n.Name == "subdir2"
	// is a file homie
	if strings.Contains(n.Name, ".ext") {
		n.IsFile = true
		return
	}

	// how many tabs does the next subdir/file need
	// if it's more than this then it's a grandchild
	tabs := count_tabs(parent) + 1

	for _, v := range s {
		// see which are children and which aren't
		ischild := (count_tabs(v) == tabs)

		if ischild {

			ss := get_substring(str, v)
			isfile := strings.HasSuffix(strings.TrimSpace(v), ".ext")
			// there are no further children, stop here.
			if !isfile && ss == v {
				continue
			}

			child := scan_string(ss)

			n.Children = append(n.Children, &child)
		}
	}

	return
}

// basically returns an array of paths to files.
func get_all_paths(n Node) (ret []string) {

	base := n.Name
	for _, v := range n.Children {
		if v.IsFile {
			return []string{path.Join(base, v.Name)}
		}

		v.Name = path.Join(base, v.Name)
		arr := get_all_paths(*v)
		if len(arr) > 0 {
			ret = append(ret, arr...)
		}
	}

	return
}

func solve(str string) (max int) {
	no := scan_string(str)
	for _, v := range get_all_paths(no) {
		length := len(strings.ReplaceAll(v, "/", ""))
		if length > max {
			max = length
		}
	}

	return
}

var cases = map[int]int{
	solve("dir\n\tsubdir1\n\tsubdir2"):               0,
	solve("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"): 18,
	solve("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\t" + "subsubdir2\n\t\t\tfile2.ext"): 29,
}

func main() {

	for k, v := range cases {
		if k != v {
			fmt.Println("something's wrong buddy")
			os.Exit(1)
		}
	}
	fmt.Println("success")

}

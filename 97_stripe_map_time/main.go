package main

import (
	"fmt"
)

type val struct {
	value uint
	time int
}

type valmap map[int]val

type timemap map[int]valmap

var globalmap = timemap{}

func set(key int, value uint, time uint) {
	v, ok := globalmap[key]
	newval := val{value, int(time)}
	if !ok {
		globalmap[key] = map[int]val{int(time): newval}
		globalmap[key][-1] = globalmap[key][int(time)]
	} else {
		latest := v[-1]
		if latest.time < int(time) {
			globalmap[key][-1] = newval
		}
		
		globalmap[key][int(time)] = newval
	}
}

func get(key int, time int) int {
	_, ok1 := globalmap[key]
	if ok1 {
		valu, ok := globalmap[key][time]
		if ok { 
			return int(valu.value)
		} else {
			if globalmap[key][-1].time < time {
				return int(globalmap[key][-1].value)
			}
			
			var last *val
			for _, v := range globalmap[key] {
				if last == nil {
					if v.time < time { last = &v }
				} else {
					if (last.time < v.time || last.time > time) && v.time < time {
						last = &v
					}
				}
			}
			
			if last != nil {
				return int(last.value)
			}
		}
	}
	
	return -1
}

func main() {
	set(1, 1, 0)
	set(1, 2, 2)
	
	fmt.Println(get(1, 1), get(1, 3))
	globalmap = timemap{}
	
	set(1, 1, 5)
	fmt.Println(get(1, 0))
	fmt.Println(get(1, 10))
	
	globalmap = timemap{}
	set(1, 1, 0)
	set(1, 2, 0)
	fmt.Println(get(1, 0))
}

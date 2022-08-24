package main

import "fmt"

type lru struct {
	values map[string]string
	keys   []string
	limit  int // if -1, then no limit
}

func newLRU(size int) *lru {
	l := &lru{
		limit:  size,
		values: map[string]string{},
		keys:   []string{},
	}

	return l
}

// O(1), space O(n*2)
func (l *lru) set(key, value string) {
	if l.limit > 0 {
		if len(l.keys) == l.limit {
			key := l.keys[0]
			l.keys = l.keys[1:]

			delete(l.values, key)
		}
	} else if l.limit == 0 {
		return
	}

	l.keys = append(l.keys, key)
	l.values[key] = value
}

// O(1), space O(n*2)
func (l *lru) get(key string) string {
	return l.values[key]
}

func main() {
	l := newLRU(4)

	l.set("first", "value1")
	l.set("second", "value2")
	l.set("third", "value3")
	l.set("fourth", "value4")

	fmt.Println("initalizing lru with size 4")
	fmt.Println("getting first element:", l.get("first"))
	l.set("fifth", "value5")
	fmt.Println("getting first element after adding fifth element:", l.get("first"))
	fmt.Println("getting second element:", l.get("second"))
	fmt.Println("getting fifth element:", l.get("fifth"))

}

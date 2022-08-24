package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type urlsh struct {
	length  int
	selchar []byte
	db      map[string]string
}

func (u *urlsh) generate() string {

	str := []byte{}
	for i := 0; i < u.length; i++ {
		j := rand.Intn(len(u.selchar))
		str = append(str, u.selchar[j])
	}

	return string(str)
}

func (u *urlsh) shorten(earl string) string {
	for k, v := range u.db {
		if v == earl {
			return k
		}
	}

	shorty := u.generate()
	u.db[shorty] = earl

	return shorty
}

func (u *urlsh) restore(short string) string {
	return u.db[short]
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	u := &urlsh{
		length:  6,
		selchar: []byte(characters),
		db:      map[string]string{},
	}

	earl := "https://lemondev.xyz"
	fif := u.shorten(earl)
	sec := u.shorten(earl)

	if sec != fif {
		fmt.Println("bad code, same url different id")
		os.Exit(1)
	}
	if u.restore(fif) != earl {
		fmt.Println("bad code, restore doesn't work.")
		os.Exit(1)
	}
	fmt.Println("success")
}

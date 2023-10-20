package main

import (
	"fmt"

	"github.com/lyonnee/hpool"
)

type People struct {
	Name string
	Age  int
}

func main() {
	pool := hpool.New[*People](func() *People {
		return &People{}
	})

	p := pool.Get()
	fmt.Println(p)

	pool.Put(p)
}

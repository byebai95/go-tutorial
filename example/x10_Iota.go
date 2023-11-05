package main

import "fmt"

const (
	a = (1 << iota) / 4
	b
	c
	d
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

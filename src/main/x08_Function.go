package main

import "fmt"

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	var a, b = swap(10, 20)
	fmt.Println(a, b)
}

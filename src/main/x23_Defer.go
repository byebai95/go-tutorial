package main

import "fmt"

// defer 执行函数时，函数执行完毕再执行 defer 栈中的内容
func sum(n, m int) int {

	defer fmt.Println("a")
	defer fmt.Println("b")

	fmt.Println("c")
	return m + n
}

func main() {

	result := sum(10, 20)
	fmt.Println("d", result)
}

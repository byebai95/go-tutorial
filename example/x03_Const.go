package main

import "fmt"

//全局常量
const (
	NAME = "张三"
	AGE  = "18"
)

func main() {
	//局部常量
	const LENGTH int = 1024

	fmt.Println(LENGTH)
	fmt.Println(NAME)
}

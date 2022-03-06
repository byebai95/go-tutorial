package main

import "fmt"


//全局匿名函数
var(
	sub = func(n,m int)int{
		return n-m
	}
)

func main() {

	//局部匿名函数
	var sum = func(n, m int) int {
		return n + m
	}
	fmt.Println(sum(10,20))

	fmt.Println(sub(10,20))
}

package main

import "fmt"

func main() {
	// 默认零值
	var bol bool
	fmt.Println(bol)

	var i int
	fmt.Println(i)

	//根据值判定类型
	var j = 100
	fmt.Println(j)

	k := 20
	fmt.Println(k)

	//值类型与引用类型
	var e = 10
	fmt.Println(&e)
}

var a, b int

//全局变量声明
var (
	c int
	d int
)

/**
变量声明
1.变量是由数字、字母、下划线组成，首个字符不能是数字
2.变量声明未初始化，默认为零值
	bool 默认 false
	int  默认 0
	字符串为 ""
	指针、数组、map、chan、函数类型等是 nil
3. := 等于申明+初始化,只能在函数体出现
4.
*/

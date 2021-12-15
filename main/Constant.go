package main

import (
	"fmt"
	"unsafe"
)

//常量枚举
const (
	MALE   = 0
	FEMALE = 1
	ADDR   = "beijing"
)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {

	const age int = 20
	//age = 30 常量不可修改
	fmt.Println(age)

	const name = "zhangsan"
	//name = "lisi"  隐式定义
	fmt.Println(name)

	//
	fmt.Println(len(ADDR), unsafe.Sizeof(ADDR))

	//iota
	//1
	fmt.Println(i, j, k, l)
}

/**
1.iota特殊常量，可以被编译器修改的常量， iota const 出现时为 0 ，每新增一行常量声明 iota 计数一次
*/

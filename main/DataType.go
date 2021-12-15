package main

import (
	"fmt"
	_ "fmt"
)

func main() {

	var b bool = true

	var age uint8 = 25

	fmt.Println(b, age)

}

/**
1.布尔类型
	var b bool = true
2.数字类型

	uint8 无符号8位整形
	uint16 无符号16位整形
	uint32
	uint64
	int8
	int32
	int64

	float32 32位浮点类型
	float64
	complex64 32位实数与虚数
	complex128

	其他类型
	byte
	rune
	uint
	int
	uintptr 无符号整形，用于存放指针
3.字符串类型
4.派生类型
	指针、数组、结构化、channel 、函数、切片、接口、map
*/

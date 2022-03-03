package main

import "fmt"

func main() {

	//整形
	var age int = 18
	fmt.Println(age)

	//字符类型
	str := "abc"
	fmt.Println(str)

	//浮点类型
	var grade float32
	grade = 80.1
	fmt.Println(grade)

	//布尔类型
	var tag bool = true
	fmt.Println(tag)

	//指针类型
	var ptr *int = &age
	fmt.Println(*ptr)

	//数组类型.
	var arr = [...]int{1, 10, 100}
	fmt.Println(arr)

	//切片类型，动态数组
	var sli []int
	fmt.Println(sli)

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

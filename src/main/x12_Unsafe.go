package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n = 12.34
	fmt.Printf("变量 n 的类型是 %T ，占用内存大小为 %d\n", n, unsafe.Sizeof(n))

	var i = 10
	fmt.Printf("变量 i 的类型是 %T ，占用内存大小为 %d\n", i, unsafe.Sizeof(i))

	var bol bool = true
	fmt.Printf("变量 bol 的类型是 %T ，占用内存大小为 %d\n", bol, unsafe.Sizeof(bol))

}

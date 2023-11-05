package main

import (
	"fmt"
	"strconv"
)

func main() {

	//var str string = "abcdef"
	//str[0] = 'p' ,无法对字符串进行单个字符修改+
	//但是可以单个字符输出

	//其他类型转换为 string
	var num int = 10
	var f float32 = 10.01
	var nStr = fmt.Sprintf("%d", num)
	var fStr = fmt.Sprintf("%f", f)
	fmt.Printf("num %d 类型： %T \n", num, nStr)
	fmt.Printf("fStr %f 类型 %T \n", f, fStr)

	//方式二
	strconv.FormatInt(int64(num), 10)
	strconv.FormatFloat(float64(f), 10, 10, 64)

	//string 转 基本类型
	var s1 string = "100"
	i1, _ := strconv.ParseInt(s1, 10, 64)
	fmt.Printf("%d = %T\n", i1, i1)

	//string 转基本失败,值为默认值
	var s2 = "3.14"
	i2, _ := strconv.ParseInt(s2, 10, 64)
	fmt.Printf("%d = %T\n", i2, i2)
}

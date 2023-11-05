package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//字符串输出中文乱码，UTF8 中文占用三个字节,转换[]rune防止乱码
	str := "Hello,北京"
	strRun := []rune(str)
	for i := 0; i < len(strRun); i++ {
		fmt.Printf("%c", strRun[i])
	}

	//string 与 数字转换
	i := 10
	iStr := strconv.Itoa(i)
	fmt.Println(iStr)

	j, err := strconv.Atoi(iStr)
	if err != nil {
		fmt.Println("转换异常", err)
	} else {
		fmt.Println("转换成功", j)
	}

	//字符串最后出现的位置
	index := strings.LastIndex("Hello,World", "o")
	fmt.Println(index)

	//字符串替换 replace
	s1 := strings.Replace("abcda", "a", "e", 2)
	fmt.Println(s1)

	//字符串分割
	strArr := strings.Split("Hello,World", ",")
	fmt.Println(strArr)

	//去除空格
	s2 := strings.TrimSpace(" Hello,World  ")
	fmt.Println(s2)

	//判断开头
	s3 := strings.HasPrefix("www.xxx.com", "www")
	fmt.Println(s3)

	//字符串替换
	str2 := "abcd"
	//str2[0] = 'o' ,报错，string 是不可修改的
	s4 := []rune(str2)
	s4[0] = 'p'
	fmt.Println(string(s4))

}

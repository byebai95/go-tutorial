package main

import "fmt"

func main() {
	var age int
	var name string
	var sex string

	fmt.Println("请输入姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入年龄：")
	fmt.Scanf("%d\n", &age)
	fmt.Println("请输入性别：")
	fmt.Scanf("%s\n", &sex)

	fmt.Printf("【个人信息】 姓名：%s  年龄：%d 性别：%s\n", name, age, sex)

	//方式二 scan
	var msg string
	fmt.Println("请输入整行信息：")
	fmt.Scanln(&msg)
	fmt.Printf("【msg】:%s", msg)
}

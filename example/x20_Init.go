package main

import "fmt"

//在 main 方法之前执行
//如果全局变量，执行顺序 全局变量- init() - main()
func init() {
	fmt.Println("init()")
}

func main() {

	fmt.Println("main 执行")
}

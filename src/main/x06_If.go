package main

import "fmt"

func main() {

	var age int = 18

	if age < 18 {
		fmt.Println("未成年")
	} else if age < 60 {
		fmt.Println("已成年")
	} else {
		fmt.Println("已退休")
	}
}

/**
1. if
2. if - else
3. if 切套
4. switch
5. select
	-类似 switch ,select 会随机执行一个可运行的 case .如果没有 case 可执行，将进行阻塞直到 case 可运行
6. go 没有三目运算符
*/

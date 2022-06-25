package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for false {
		fmt.Println("循环")
	}

	var str = "Hello,World"

	for index, value := range str {
		fmt.Printf("index = %d, value = %c\n", index, value)
	}

	//输出 99乘法表

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j <= i {
				fmt.Printf("%d*%d=%d ", j, i, i*j)
			}
		}
		fmt.Println()
	}

}

/**
1. for 循环
2. 循环嵌套
3.循环控制语句
	- break
	- continue
	- goto 将循环转移到被标记的语句
*/

package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	for false {
		fmt.Println("循环")
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

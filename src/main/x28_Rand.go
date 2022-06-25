package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//防止每次生成重复数据
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(100), " ")
	}
}

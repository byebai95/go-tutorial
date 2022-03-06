package main

import (
	"fmt"
	"time"
)

func main(){

	//当前时间
	now := time.Now()
	fmt.Println(now)

	//格式化输出
	fmt.Println(now.Format("0000/10"))

	//睡眠
	time.Sleep(time.Second * 1)

	//时间戳
	fmt.Println(now.UnixNano())

	//计算执行时间
	now.Unix()
	now.Unix()
}

package main

import (
	"fmt"
	"io/ioutil"
)

func main(){


	byteArr,err := ioutil.ReadFile("d:/test.txt")

	if err != nil{
		fmt.Println("打开异常")
	}

	fmt.Println(string(byteArr))

}

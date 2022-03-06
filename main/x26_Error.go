package main

import "fmt"

func calc(){

	//必须定义在首部，否则无法捕获异常
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("抛出异常：",err)
		}
	}()

	n := 10
	m := 0
	res := n/m
	fmt.Println("计算结果：",res)

}

func main(){

	calc()

	fmt.Println("main 执行结束")
}

package main

import "fmt"


func main(){
	var arr = [3]int{1,2,3}

	change(&arr)

	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}
}


//数组默认是值传递，引用传递需要传地址
func change(arr *[3]int){
	for i:=0;i<len(arr);i++{
		arr[i] *= 10
	}
}
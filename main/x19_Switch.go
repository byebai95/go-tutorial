package main

import "fmt"

func main(){

	var score float32
	fmt.Println("请输入分数：")
	fmt.Scanf("%f",&score)

	switch {
	case score < 60 :{
		fmt.Println("不及格")
	}
	case score >=60 && score <90 :{
		fmt.Println("优秀")
	}
	case score >=90 :{
		fmt.Println("顶级分数")
	}
	default:
		fmt.Println("非法输入")
	}


}

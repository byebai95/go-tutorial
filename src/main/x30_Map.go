package main

import "fmt"

func main(){


	var city = map[string]string{
		"beijing":"北京",
		"shanghai":"上海",
	}


	//删除某个元素
	delete(city,"beijing")


	//查找某个元素
	value,bol := city["shanghai"]
	if bol {
		fmt.Println("查找到 value ",value)
	}else{
		fmt.Println("未找到")
	}
	fmt.Println(city)

}

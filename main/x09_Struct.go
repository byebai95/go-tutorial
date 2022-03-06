package main

import "fmt"

type People struct {
	name string
	age int
	sex string
	address string
}


func main(){

	p1 := People{"张三",15,"男","北京"}

	fmt.Println(p1)

	var p2 People
	p2.name = "李四"
	fmt.Println(p2)

	var p3 *People
	p3 = new(People)
	(*p3).name = "王五"
	fmt.Println(*p3)

}

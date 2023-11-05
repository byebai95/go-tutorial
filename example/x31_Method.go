package main

import "fmt"

type Dog struct {
	age int
}

//方法，只能被绑定的类型调用
func (dog Dog) eat() {
	fmt.Println("【dog】:", dog.age)
}

func main() {

	dog := Dog{21}

	dog.eat()

}

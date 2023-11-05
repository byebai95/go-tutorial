package main

import "fmt"

type Animal interface {
	eat()
	sleep()
}

type Cat struct {
	name string
}

type Human struct {
	name string
}

func (cat Cat) eat() {
	fmt.Println(cat.name, "is eating")
}

func (cat Cat) sleep() {
	fmt.Println("")
}

func (human Human) eat() {
	fmt.Println(human.name, "is eating")
}

func main() {

	//子结构必须全部实现接口的方法
	var animal Animal
	animal = Cat{"小猫"}
	animal.eat()

}

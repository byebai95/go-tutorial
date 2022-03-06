package main

import "fmt"

//闭包

func Add() func(n int)int{
	m := 10
	return func(n int) int{
		m += n
		return m
	}
}

func main(){

	f := Add()

	fmt.Println(f(1))
	fmt.Println(f(2))


}

package main

import "fmt"

func main() {

	arr := [5]int{3, 2, 5, 8, 4}

	slice := arr[2:4]

	fmt.Println(slice)

	fmt.Println("lenSlice : ", len(slice))

	fmt.Println("capSlice : ", cap(slice))

	fmt.Println("&slice[0] : ", &slice[0])
	fmt.Println("&arr[2] : ", &arr[2])

	//切片创建,切片是指向数组的指针
	s1 := make([]int, 2, 4)
	fmt.Println("s1", s1)

	//append 测试
	s1 = append(s1, 10)
	fmt.Println("s1", s1)

	//切片内部指向匿名数组，测试
	s2 := make([]int, 3, 5)
	var ptr *int
	ptr = &s2[0]
	fmt.Println(ptr)

	str := "HelloWorld"
	s3 := str[2:6]
	fmt.Println(s3)

}

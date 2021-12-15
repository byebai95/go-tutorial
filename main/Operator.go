package main

import "fmt"

func main() {

	var i = 100
	fmt.Println(&i)
	fmt.Println(*&i)

}

/**
1.内置运算符
	-算数运算符
	-关系运算符
	-逻辑运算符
	-位运算符
	-赋值运算符与
	-其他运算符
2.算数运算符
	+
	-
	*
	/
	%
	++
	--
3.关系运算符
	==
	!=
	>
	>=
	<
	<=
4.逻辑运算符
	&&
	||
	！
5.位运算符
	& 按位与
	|
	^ 异或
	<< 按位左移，高位丢弃，低位补 0
	>> 按位右移，高位补0，低位丢弃
6.赋值运算符
	=
	+=
	-=
	*=
	/=
	%=
	<<=
	>>=
	&=
	^=
	|=
7.其他运算符
	& 返回变量存储地址
	* 指针变量
*/

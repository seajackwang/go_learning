package main

import "fmt"

func main() {
	/*
		运算符
		1.算术运算  + - * / ++ --
		2.关系运算符 == >= <= != < >
		3.逻辑运算 && || !
		4.位运算 & | ^
		5.赋值运算 =  += -= *= /=  ...
	*/
	var a, b = 1, 4
	c := a + b
	d := a - b
	var f = a * b
	var g = a / b
	a++
	fmt.Println(a)
	a = 21
	a--
	fmt.Println(a)

	fmt.Println(c, d, f, g)
	var rr = 0
	for a < 30 {
		a++
		rr += a
		fmt.Println(rr)
	}
}

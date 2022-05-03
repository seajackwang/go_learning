package main

import "fmt"

var kk, jj, ll float64 = 1.2, 2.3, 4.8

func main() {
	//数据类型
	/*
		1.布尔型
		2.数字类型 --整形int（uint32-int32;有没有符号） 浮点型float32、float64 复数 中位数
		3.字符串 go使用UTF-8编码
		4.派生类型
			--指针类型
			--数组
			--结构化（struct)
			--Channel类型
			--函数
			--切片类型
			--接口（interface）
			--Map类型
	*/
	var b = "RUNoob"
	fmt.Println(b)
	var a int16
	var c bool
	fmt.Println(a, c)

	var p *int
	var mm []int
	var e error

	fmt.Println(p, mm, e)
	fmt.Println(kk, jj, ll)
	_, nub, strs := nubers()
	fmt.Println(nub, strs)

}

func nubers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}

package main

import "fmt"

const (
	unknow = 0
	male   = 1
	female = 2
)

func main() {
	fmt.Println("你好")
	var a int = 1
	a = 100
	b := 3
	const c = 4
	const d string = "abc"

	res := (a + b) / c
	fmt.Println(a + b)
	fmt.Println(res)
	fmt.Println(d)
	fmt.Println(unknow, male, female)

	const (
		ss = len(d)
		sf = c + c
	)
	fmt.Println(ss, sf)

	const (
		g = iota
		h
		j
		k
		l = "h"
		z
		n = 100
		m = iota
	)
	fmt.Println(g, h, j, k, l, z, n, m)

}

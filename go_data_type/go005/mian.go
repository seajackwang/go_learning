package main

import "fmt"

func main() {
	var d,f = 10,13
	var c int
	c = max(d,f)
	fmt.Println(c)
}

func max(num1, num2 int) int {
	var result int
	if (num1 > num2){
		result = num1
	} else {
		result = num2
	}
	return result

}
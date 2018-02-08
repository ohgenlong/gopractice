package main

import (

	"fmt"
)

// 闭包
func Adder() func(int) int {
	var x int
	return func(a int) int {
		x += a
		return x
	}
}

func main() {
	f := Adder()
	fmt.Println(f(1))
	fmt.Println(f(10))
	fmt.Println(f(100))
}
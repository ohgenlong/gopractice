package main

import "fmt"

var a int

func main() {
	a = 10
	fmt.Println(a)

	f1()
}

func f1() {
	a := 20
	fmt.Println(a)
	f2()
}

func f2() {
	fmt.Println(a)
}

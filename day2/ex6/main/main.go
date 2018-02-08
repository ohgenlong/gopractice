package main

import "fmt"

var a int = 10

func n() {
	fmt.Println(a)
}

func m() {
	a = 20
	fmt.Println(a)
}
func main() {
	n()
	n()
	m()
}

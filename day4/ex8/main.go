package main

import "fmt"

var b [10]int

func test2() {
	var a [10]int
	b := a
	b[0] = 100
	fmt.Println(b)
}

func test3(ab [5]int) {
	ab[0] = 1001
}

func test4() {
	b[0] = 1000
}

func test5(ab *[10]int) {
	ab[0] = 111
}

func main() {
	var a [5]int
	//var b [20]int
	a[0] = 100
	//a[10] = 200
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	for index, val := range a {
		fmt.Println(index, val)
	}

	test2()
	//var b [5]int
	test3(a)
	test4()
	fmt.Println(b)

	var bb [10]int
	test5(&bb)
	fmt.Println(bb)
}

package main

import "fmt"

func set(p *int) {
	*p = 123
}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("catch panic: %v\n", err)
		}
	}()

	var p *int
	set(p)
	fmt.Printf("*p:%d\n", *p)
}

func main() {
	test()
}

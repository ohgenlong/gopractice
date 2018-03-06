package main

import "fmt"

func swap1(a *int, b *int) {
	fmt.Println("ori a, b:")
	fmt.Println(&a, &b)
	tmp := *a
	*a = *b
	*b = tmp
	fmt.Println("a, b:")
	fmt.Println(&a, &b)
	return
}

func swap2(a int, b int) {
	fmt.Println("ori a, b:")
	fmt.Println(&a, &b)
	tmp := a
	a = b
	b = tmp
	fmt.Println("a, b:")
	fmt.Println(&a, &b)
	return
}

func main() {
	c := 100
	d := 200
	fmt.Println("ori")
	fmt.Println(&c, &d)
	swap1(&c, &d)
	swap2(c, d)
	fmt.Println("last")
	fmt.Println(&c, &d)
	fmt.Println(c, d)
}

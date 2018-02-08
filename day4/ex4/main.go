package main

import "fmt"

func calc(n int) int {
	if n == 1 {
		return 1
	}

	return calc(n-1) * n
}

func main() {
	n := calc(5)
	fmt.Println(n)
}

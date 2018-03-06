package main

import (
	"fmt"
)

func plus(a int, b int) {
	fmt.Printf("%d+%d=%d\n", a, b, a+b)
}

func main() {
	n := 5
	for i := 0; i <= n; i++ {
		plus(i, n-i)
	}
}

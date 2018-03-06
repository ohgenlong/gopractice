package main

import (
	"fmt"
	"time"
)

func recusive(n int) {
	fmt.Println("hee")
	time.Sleep(time.Millisecond * 500)
	if n > 10 {
		return
	}
	recusive(n + 1)
}

func main() {
	recusive(1)
}

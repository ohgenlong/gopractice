package main

import (
	"time"
	"fmt"
)

const (
	Man = 1
	Female = 2
)

func main() {
	for {
		s := time.Now().Unix()
		if (s % Female == 0) {
			fmt.Println("female: ", s)
		} else {
			fmt.Println("Man: ", s)
		}
		time.Sleep(time.Millisecond*100)
	}

}
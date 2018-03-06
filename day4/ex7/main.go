package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}

func main() {
	f1 := makeSuffix(".bmp")
	f2 := makeSuffix(".jpg")
	fmt.Println(f1("test1"))
	fmt.Println(f1("test2"))
	fmt.Println(f2("test3"))
	fmt.Println(f2("test4"))
}

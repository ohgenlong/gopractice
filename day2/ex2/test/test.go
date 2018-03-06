package test

import "fmt"

var Name string = "test"
var Age int = 200

func init() {
	fmt.Println("test")
	fmt.Println("test Name: ", Name)
	fmt.Println("test Age: ", Age)

	Age = 170
	fmt.Println("test Age: ", Age)

}

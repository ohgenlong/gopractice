package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func Test(a interface{}) {
	b, ok := a.(Student)
	if ok == false {
		fmt.Println("convert failed")
		return
	}

	fmt.Println(b)
}

func Justy(items ...interface{}) {
	for _, v := range items {
		switch v.(type) {
		case string:
			fmt.Println("string: ", v)
		case int:
			fmt.Println("int: ", v)
		case bool:
			fmt.Println("bool: ", v)
		case float32, float64:
			fmt.Println("float32, float64: ", v)
		case Student:
			fmt.Println("Struct: ", v)
		case *Student:
			fmt.Println("*Struct: ", v)
		}
	}
}

func main() {
	var b Student
	Test(b)
	b.Name = "abc"
	b.Age = 123

	Justy(b, &b, 32, "abc", true, 3.24434)
}

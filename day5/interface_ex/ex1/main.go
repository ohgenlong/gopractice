package main

import "fmt"

type T_int interface {
	Print()
}

type Student struct {
	name string
	age  int
}

func (this *Student) Print() {
	fmt.Println("name: ", this.name)
	fmt.Println("age: ", this.age)
}

func main() {
	var t T_int
	var stu Student = Student{
		name: "stu1",
		age:  18,
	}

	t = &stu
	t.Print()
}

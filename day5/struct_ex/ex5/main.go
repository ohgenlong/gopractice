package main

import "fmt"

type Student struct {
	Name string
	Age int
}

type interger int


func (this *interger) print() {
	fmt.Println("p is: ", *this)
}

func (this *interger) set(v interger) {
	*this = v
}

func (this *Student) init(name string, age int) {
	this.Name = name
	this.Age = age
	fmt.Println(this)
}


func (this *Student) get() *Student {
	return this
}


func main() {
	var stu Student
	(&stu).init("stu01", 19)  // 下面的一行是一样意思，自动转了
	stu.init("stu01", 19)
	fmt.Println(stu.Name)
	fmt.Println(stu.Age)

	stu1 := stu.get()
	fmt.Println(stu1)

	var a interger
	a = 100
	a.print()

	a.set(1000)
	a.print()
}

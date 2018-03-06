package main

import "fmt"

type Car interface {
	GetName() string
	Run()
	DiDi()
}

type BMW struct {
	Name string
}

func (this BMW) GetName() string {
	//fmt.Printf("%s", this.Name)
	return this.Name
}

func (this *BMW) DiDi() {
	fmt.Printf("%s didi\n", this.Name)
}

func (this *BMW) Run() {
	fmt.Printf("%s running\n", this.Name)
}

func main() {
	var a interface{}
	var b int
	a = b
	fmt.Printf("type of a is %T\n", a)

	var car Car
	var bmw *BMW = &BMW{
		Name: "BMW",
	}

	fmt.Printf("type of car is %T\n", car)

	car = bmw // car是interface类型，默认是指针（内存地址），bmw也是指针类型才能赋值
	car.DiDi()
	car.Run()
}

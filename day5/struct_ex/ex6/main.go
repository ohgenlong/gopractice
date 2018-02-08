package main

import "fmt"

type Car struct {
	weight int
	name string
}

type Bike struct {
	Car
	oo int
}

type Train struct {
	c Car
	oo int
}

type minCar struct {
	Car
}

func (this *Car) Run() {
	fmt.Println("running")
}

func (this *minCar) String() string {
	str := fmt.Sprintf("name=[%s], weigth=[%d]", this.name, this.weight)
	return str
}


func main() {
	var a Bike

	a.weight = 100
	a.name = "bike"
	a.oo = 2
	a.Run()

	fmt.Println(a)

	var b Train
	b.c.weight = 100
	b.c.name = "train"
	b.oo = 20
	b.c.Run()

	fmt.Println(b)

	var cc minCar
	cc.weight = 60
	cc.name = "miniCar"

	fmt.Println(&cc)

}

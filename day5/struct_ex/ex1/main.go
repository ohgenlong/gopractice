package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int32
	score float32
}

func main() {
	var stu Student
	stu.Age = 18
	stu.Name = "liming"
	stu.score = 80

	var stu1 *Student = &Student{
		Age:   20,
		Name:  "kate",
		score: 100,
	}

	fmt.Println(stu)
	fmt.Println(stu.Name)
	fmt.Println(stu.Age)
	fmt.Println(stu.score)
	fmt.Println(&stu.Name)
	fmt.Println(&stu.Age)
	fmt.Println(&stu.score)

	fmt.Println(stu1)
	fmt.Println(*stu1)
	fmt.Println(stu1.Name)
	fmt.Println(stu1.Age)
	fmt.Println(stu1.score)

	fmt.Println(&stu1.Name)
	fmt.Println(&stu1.Age)
	fmt.Println(&stu1.score)
}

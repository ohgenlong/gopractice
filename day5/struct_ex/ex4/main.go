package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Student struct {
	Name string `json:"stu_name"`
	Age  int    `json:"age"`
}

type Car struct {
	name string
	age  int
}

type Train struct {
	Car
	int
	start time.Time
}

func main() {
	var stu Student = Student{
		Name: "stu01",
		Age:  16,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode stu failed")
		return
	}

	fmt.Println(string(data))

	var t Train
	t.name = "train"
	t.age = 100
	t.int = 200

	fmt.Println(t)
}

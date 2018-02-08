package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name string
	Id string
	Age int
}

type StudentArray []Student


func (this StudentArray) Len() int {
	return len(this)
}

func (this StudentArray) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this StudentArray) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	var stus StudentArray

	for i:=0 ; i<= 10; i++ {
		stu := Student {
			Name: fmt.Sprintf("stu%d", rand.Intn(i)),
			Id: fmt.Sprintf("110%d", rand.Int()),
			Age: rand.Intn(100),
		}
		stus = append(stus, stu)
	}

	for _,v := range stus {
		fmt.Println(v)
	}

	sort.Sort(stus)

	fmt.Println("#############")
	for _,v := range stus {
		fmt.Println(v)
	}
}
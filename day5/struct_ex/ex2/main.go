package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name string
	Age  int
	next *Student
}

func Trace(p *Student) {
	fmt.Println("p: ", &p)
	for p != nil {
		fmt.Println(*p)
		p = (*p).next
	}
}

func BackTrace(p *Student) {
	for &p != nil {
		fmt.Println(*&p)
		p = (*p).next
	}
}

func InsertTail(tail *Student) {
	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%d", i),
			Age:  rand.Intn(100),
		}
		tail.next = &stu
		tail = &stu
	}
}

func InsertHead1(head *Student) {
	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%d", i),
			Age:  rand.Intn(100),
		}

		stu.next = head
		head = &stu
	}
}

func InsertHead2(head **Student) {
	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%d", i),
			Age:  rand.Intn(100),
		}
		stu.next = *head
		*head = &stu
	}
}

//func DelNode(p **Student) {
//
//	var prev **Student = p
//	for *p != nil {
//		if p.Name == "stu6" {
//			*&prev.next = *&p.next
//			break
//		}
//		*prev = *p
//		*p = *&p.next
//	}
//}

func AddNode(p *Student, newnode *Student) {
	for p != nil {
		if p.Name == "stu5" {
			newnode.next = p.next
			p.next = newnode
			break
		}
		p.next = p
	}
}

func main() {

	var head *Student = new(Student)
	head.Name = "head"
	head.Age = 1

	//InsertTail(head)
	//InsertHead1(head)
	InsertHead2(&head)
	//DelNode(&head)
	Trace(head)

}

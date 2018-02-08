package main

import "fmt"

type Student struct {
	Name string
	left *Student
	right *Student
}

func transDeep(root *Student) {
	// 深度优先
	if root == nil {
		return
	}

	fmt.Println(root)
	transDeep(root.left)
	transDeep(root.right)
}



func main() {
	var root *Student = new(Student)
	root.Name = "stu1"

	var left1 *Student = new(Student)
	left1.Name = "stu2"

	root.left = left1

	var right1 *Student = new(Student)
	right1.Name = "stu4"

	root.right = right1

	var left2 *Student = new(Student)
	left2.Name = "stu3"

	left1.left = left2

	transDeep(root)
}

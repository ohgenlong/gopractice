package main

import (
	"fmt"
	"math/rand"
	"encoding/json"
	"os"
	"io/ioutil"
)

type Student struct {
	Name string `json:"name"`
	Age int	`json:"age"`
}


func jsonWriteFile(filename string) {
	studentStruct := make([]Student, 0)
	for i:=1; i< 1000; i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%d", i),
			Age: rand.Intn(20),
		}
		studentStruct = append(studentStruct, stu)
	}

	data, err := json.Marshal(&studentStruct)
	if err != nil {
		fmt.Printf("json dump err: %s\n", err)
		return
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		fmt.Printf("create file err: %s\n", err)
		return
	}

	n, err := file.Write(data)
	if err != nil {
		fmt.Printf("write file err: %s\n", err)
		return
	}

	fmt.Printf("wirte %d success\n", n)
}

func jsonReadFile(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		fmt.Printf("create file err: %s\n", err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("read file err: %s\n", err)
		return
	}

	var stus []*Student
	err = json.Unmarshal(data, &stus)
	if err != nil {
		fmt.Printf("json load file err: %s\n", err)
		return
	}

	fmt.Printf("total student: %d\n", len(stus))
	for i:=0; i<10; i++ {
		fmt.Printf("stu: %s, age: %d", stus[i].Name, stus[i].Age)
	}


}


func main() {

	//jsonWriteFile("jsonTest.json")
	jsonReadFile("jsonTest.json")
}

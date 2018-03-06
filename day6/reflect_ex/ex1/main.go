package main

import (
	"fmt"
	"reflect"
)

type ss struct {
	Name string
	Age  int
}

func test(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println(t)

	v := reflect.ValueOf(b)
	fmt.Println(v)

	k := v.Kind()
	fmt.Println(k)

	iv := v.Interface()
	stu, ok := iv.(ss)
	if ok {
		fmt.Printf("%v %T %s\n", stu, stu, reflect.TypeOf(stu).Kind())
	} else {
		fmt.Printf("%v %v\n", iv, reflect.TypeOf(iv).Kind())
	}
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	c := val.Int()
	fmt.Printf("get value interface{} %d\n", c)
}

func TestSetInt(b interface{}) {
	val := reflect.ValueOf(b)
	val.Elem().SetInt(200)
	fmt.Printf("get value interface{} %d\n", val.Elem().Int())
}

func main() {
	var a = 100
	test(a)

	fmt.Println("################")
	var b ss = ss{
		Name: "b",
		Age:  123,
	}
	test(b)
	fmt.Println("################")
	var c = "100000"
	test(c)

	fmt.Println("###############")
	testInt(1234)

	d := 2
	TestSetInt(&d)
	fmt.Println(d)

}

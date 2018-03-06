package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"stu_name"`
	Age  int    `json:"stu_age"`
}

func (this Student) TestFunc() {
	fmt.Println("Student func test", this)
}

func (this *Student) PtrTestFunc() {
	fmt.Println("Student func test")
}

func TestStruct(a interface{}) {
	v := reflect.ValueOf(a)
	ty := reflect.TypeOf(a)
	kd := v.Kind()
	switch {
	case kd == reflect.Struct:
		fmt.Println("expert struct")
		num := v.NumField()
		fmt.Printf("nums of %d struct\n", num)
		if num >= 1 {
			for i := 0; i < num; i++ {
				fmt.Printf("%d %v\n", i, v.Field(i))
			}
		}
		numFunc := v.NumMethod()
		fmt.Printf("nums of %d func struct\n", numFunc)
		if numFunc >= 1 {
			//var pa []reflect.Value
			v.Method(0).Call(nil)
		}

		tag := ty.Field(0).Tag.Get("json")
		fmt.Println("tag: ", tag)

	case kd == reflect.Ptr && v.Elem().Kind() == reflect.Struct:

		num := v.Elem().NumField()
		fmt.Printf("num of %d struct\n", num)
		numFunc := v.Elem().NumMethod()
		fmt.Printf("num of %d func struct\n", numFunc)

	case kd == reflect.Int:
		fmt.Println("expert int")

	default:
		fmt.Println("not match format")
	}
}

func main() {

	var a interface{} = Student{
		Name: "stu01",
		Age:  18,
	}

	TestStruct(a)
	TestStruct(123)
	TestStruct("123")

	result, _ := json.Marshal(a)

	TestStruct(result)
}

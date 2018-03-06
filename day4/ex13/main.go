package main

import (
	"fmt"
)

func testMap() {

	var c map[string]string // 只声明不分配空间，不能直接用，需要分配内存空间

	c = make(map[string]string)

	var b map[string]string = make(map[string]string) // 声明 + 分配内存空间

	var a map[string]string = map[string]string{ // 声明 + 分配内存空间 + 初始化
		"abcd": "dddd",
	}
	a["abc"] = "abcd"
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
}

func testMap2() {
	aa := make(map[string]map[string]string, 100) // 提前分配100个长度,效率高
	aa["key1"] = make(map[string]string)
	aa["key1"]["key1"] = "1"
	aa["key1"]["key2"] = "2"
	aa["key1"]["key3"] = "3"
	aa["key1"]["key4"] = "4"
	fmt.Println(aa)
}

func testMap3() {
	aa := make(map[string]map[string]string, 100) // 提前分配100个长度,效率高
	aa["key1"] = make(map[string]string)
	aa["key1"]["key1"] = "1"
	aa["key1"]["key2"] = "2"
	aa["key1"]["key3"] = "3"
	aa["key1"]["key4"] = "4"
	_, ok := aa["key1"] // 判断是否有key1
	if ok {
		fmt.Println(aa["key1"])
	} else {
		fmt.Println("not key key1")
	}

}

func testMap4() {
	aa := make(map[string]map[string]string, 100) // 提前分配100个长度,效率高
	aa["key1"] = make(map[string]string)
	aa["key1"]["key1"] = "1"
	aa["key1"]["key2"] = "2"
	aa["key1"]["key3"] = "3"
	aa["key1"]["key4"] = "4"
	modify(&aa)

}

func modify(aaa *map[string]map[string]string) {
	aa := *aaa
	_, ok := aa["key1"]
	if ok {
		aa["key1"]["key1"] = "11" // 修改map里面的值
		fmt.Println(aa["key1"])
	} else {
		fmt.Println("not key key1")
		aa["key1"]["key1"] = "111"
	}

	for k, v := range *aaa {
		fmt.Println(k)
		for k1, v1 := range v {
			fmt.Println(k1)
			fmt.Println(v1)
		}
	}

}

func testMap5() {
	aa := make(map[string]map[string]string, 100) // 提前分配100个长度,效率高
	aa["key1"] = make(map[string]string)
	aa["key1"]["key1"] = "1"
	aa["key1"]["key2"] = "2"
	aa["key1"]["key3"] = "3"
	aa["key1"]["key4"] = "4"

	aa["key2"] = make(map[string]string)
	aa["key2"]["key1"] = "1"
	aa["key2"]["key2"] = "2"

	fmt.Println(aa)
	delete(aa, "key1")
	fmt.Println(aa)
	fmt.Println(len(aa))
}

func testMap6() {
	var a []map[int]int
	a = make([]map[int]int, 6)
	if a[0] == nil {
		a[0] = make(map[int]int)
	}
	a[0][10] = 10
	fmt.Println(a)

}

func main() {
	testMap()
	testMap2()
	testMap3()
	testMap4()
	testMap5()
	testMap6()
}

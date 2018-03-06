package main

import (
	"errors"
	"fmt"
	"time"
)

func initCofig() (err error) {
	return errors.New("init config failed")
}

func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//b := 0
	//a := 100 / b
	//fmt.Println(a)

	err := initCofig()
	if err != nil {
		panic(err)
	}
	return
}

func main() {

	for {
		test()
		time.Sleep(time.Second * 2)
	}

	var i int
	fmt.Println(i)

	j := new(int)
	*j = 100
	fmt.Println(*j)

	var a []int
	a = append(a, 10, 20, 30)
	a = append(a, a...)
	fmt.Println(a)

}

package main

import "fmt"

func testSlice() {
	var a [5]int = [...]int{1, 2, 3, 4, 5}
	s := a[1:]
	s[2] = 10
	a[2] = 4
	fmt.Printf("&a[1]: %p\n", &a[1])
	fmt.Printf("s: %p\n", s)
	fmt.Println(s)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	fmt.Printf("%d %d\n", len(s), cap(s))
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s[1] = 1000
	fmt.Printf("s: %p\n", s)
	fmt.Println(s)
	fmt.Println(a)
}

func testString() {
	s := "hello"
	s1 := []rune(s)
	s1[1] = '0'
	st := string(s1)
	fmt.Println(s1)
	fmt.Println(st)
}

func testCopy() {
	var a []int = []int{1, 1, 1, 1, 5}

	b := make([]int, 10)

	copy(b, a)
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	testSlice()
	testCopy()
	testString()
}

package main

import "fmt"

func reverse(s string) string {
	var res string
	strlen := len(s)
	for i := 0; i < strlen; i++ {
		res = res + fmt.Sprintf("%c", s[strlen-i-1])
	}
	return res
}

func reverse1(s string) []byte {
	var res []byte
	tmp := []byte(s)
	length := len(s)
	for i := 0; i < length; i++ {
		res = append(res, tmp[length-i-1])
	}
	return res
}

func main() {
	a := reverse("abcd")
	fmt.Println(a)
	b := reverse1("hello world")
	fmt.Println(string(b))
}

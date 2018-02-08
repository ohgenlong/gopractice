package main

import (
	"sort"
	"fmt"
)

func testSortInt() {
	var a = [...]int{1,2342,23,4532}
	sort.Ints(a[:])
	fmt.Println(a)
}


func testSortString() {
	var a = [...]string{"abc","sdsdff","edfwsdf","dedrf"}
	sort.Strings(a[:])
	fmt.Println(a)
}

func testSortFloat() {
	var a = [...]float64{1.123,5.123,1.23234,4.2323}
	sort.Float64s(a[:])
	fmt.Println(a)
}

func testIntSearch() {
	var a = [...]int{1,2,3,4,5,6}
	index := sort.SearchInts(a[:], 2)
	fmt.Println(index)
}

func main() {
	testSortInt()
	testSortString()
	testSortFloat()
	testIntSearch()
}
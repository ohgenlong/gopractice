package main

func fab(n int)  int{

	if n <= 1 {
		return n
	}

	return fab(n - 1) + fab(n -2)
}


func main() {
	var n int = 10
	for i:=1; i<= n; i++ {
		a := fab(i)
		println(a)
	}

}


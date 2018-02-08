package main

import (
	"oldboy/day7/homework/ex1/balance"
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {
	inSs := make([]*balance.Instance, 0)
	for i:=0; i<5; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		port := 80
		one := balance.InitInstance(host, port)
		inSs = append(inSs, one)
	}
	//var bal balance.BaLance
	//
	//conf := "random"
	//
	//if len(os.Args) > 1 {
	//	conf = os.Args[1]
	//}
	//
	//if conf == "random" {
	//	bal = &balance.RandomBanlance{}
	//	fmt.Printf("use random\n")
	//} else if conf == "roundrobin" {
	//	bal = &balance.RoundRobinBalance{}
	//	fmt.Printf("use roundrobin\n")
	//}

	var balanceName string = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}
	fmt.Printf("use balance: %s\n", balanceName)

	for {
		ins, err := balance.DoBalance(balanceName, inSs)
		if err != nil {
			fmt.Printf("balance err: %s\n", err)
			continue
		}
		fmt.Printf("%s\n", ins)
		time.Sleep(time.Second * 1)
	}


}

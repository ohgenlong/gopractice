package main

import (
	"flag"
	"fmt"
)

var (
	conf string
	level int
)

func init() {
	flag.StringVar(&conf, "c", "D:/etc/common.conf","conf file is required")
	flag.IntVar(&level, "l", 8, "log level is required")
	flag.Parse()
}

func main() {
	fmt.Printf("conf is : %s \n", conf)
	fmt.Printf("log level is : %d \n", level)
}

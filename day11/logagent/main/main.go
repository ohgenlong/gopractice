package main


import (
	"fmt"
)


func main() {
	filename := "./conf/logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err: %s\n", err)
		panic("load conf failed")
	}
	
	err = initLogger()
	if err != nil {
		fmt.Printf("init conf failed, err: %s\n", err)
		panic("init conf failed")
	}
	return
}
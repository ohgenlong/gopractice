package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "./logagent.conf")
	if err != nil {
		fmt.Println("new config failed, err: ", err)
		return
	}

	port, err := conf.Int("server::listen_port")
	if err != nil {
		fmt.Println("read server::port failed, err: ", err)
		return
	}

	fmt.Println("Port: ", port)
	log_level := conf.String("log::log_level")
	if len(log_level) == 0 {
		log_level = "debug"
	}
	fmt.Println("log level: ", log_level)

	log_path := conf.String("logs::log_path")
	fmt.Println("log path: ", log_path)

	
}
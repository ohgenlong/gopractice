package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	F "gopractice/day11/logagent/fun"
	G "gopractice/day11/logagent/g"
	Log "gopractice/day11/logagent/logs"
	T "gopractice/day11/logagent/tailf"
	// "time"
)

func main() {
	filename := "./conf/logagent.conf"
	err := G.LoadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err: %s\n", err)
		panic("load conf failed")
	}

	err = Log.InitLogger()
	if err != nil {
		fmt.Printf("init conf failed, err: %s\n", err)
		panic("init conf failed")
	}

	logs.Debug("initialize log success")
	logs.Debug("load conf succ, config: %v", G.AgentConf)

	err = T.InitTail(G.AgentConf.CollectConfSlice, G.AgentConf.ChanSize)
	if err != nil {
		logs.Error("init tail failed, err: %v", err)
		return
	}

	logs.Debug("initialize all success")
	// go func() {
	// 	var count int
	// 	for {
	// 		count++
	// 		logs.Debug("test for logger %d", count)
	// 		time.Sleep(time.Millisecond * 10)

	// 	}
	// }()

	F.ServerRun()
	// if err != nil {
	// 	logs.Error("server run failed, err: %v", err)
	// 	return
	// }

	logs.Info("program exited")
	return
}

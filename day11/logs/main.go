package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "./logcollect.log"
	// config["level"] = logs.LevelInfo
	config["level"] = logs.LevelDebug
	
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err: ", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))

	logs.Debug("this is a test, my name is %s", "stu01")
	logs.Trace("this is a test, my name is %s", "stu02")
	logs.Warn("this is a test, my name is %s", "stu03")
	logs.Critical("this is a test, my name is %s", "stu04")

}
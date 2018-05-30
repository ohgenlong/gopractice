package logs

import (
	"github.com/astaxie/beego/logs"
	"fmt"
	"encoding/json"
	"gopractice/day11/logagent/g"
)

func convertLogLevel(level string) int {
	switch (level) {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace

	}
	return logs.LevelDebug
}

func InitLogger() (err error){
	configs := make(map[string]interface{})
	configs["filename"] = g.AgentConf.LogPath
	// configs["level"] = logs.LevelInfo
	configs["level"] = convertLogLevel(g.AgentConf.LogLevel)
	
	configStr, err := json.Marshal(configs)
	if err != nil {
		fmt.Println("initLogger failed, marshal err: ", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}
package main

import (
	"fmt"
	"errors"
	"github.com/astaxie/beego/config"
)

var (
	agentConf *Config
)

type Config struct {
	logLevel string
	logPath string
	collectConf []CollectConf
}

type CollectConf struct {
	logPath string
	topic string
}

func loadCollectConf(conf config.Configer) (err error) {
	var cc CollectConf
	cc.logPath = conf.String("collect::log_path")
	if len(cc.logPath) == 0 {
		err = errors.New("invalid collect:::log_path")
		return
	}

	cc.topic = conf.String("collect::topic")
	if len(cc.topic) == 0 {
		err = errors.New("invalid collect:::topic")
		return
	}
	
	agentConf.collectConf = append(agentConf.collectConf, cc)
	return
}

func loadConf(confType string, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Printf("new config failed, err: %s", err)
		return
	}
	agentConf = &Config{}
	agentConf.logLevel = conf.String("logs::log_level")
	if len(agentConf.logLevel) == 0 {
		agentConf.logLevel = "debug"
	}

	agentConf.logPath = conf.String("logs::log_path")
	if len(agentConf.logPath) == 0 {
		agentConf.logPath = "./logs"
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed, err: %s", err)
		return
	}
	return
}


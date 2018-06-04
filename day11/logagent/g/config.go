package g

import (
	"errors"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

var (
	AgentConf *Config
)

type Config struct {
	LogLevel string
	LogPath string
	CollectConfSlice []CollectConf
	ChanSize int
	KafkaAddr string
}

type CollectConf struct {
	LogPath string
	Topic string
}

func loadCollectConf(conf config.Configer) (err error) {
	var cc CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect:::log_path")
		return
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect:::topic")
		return
	}
	
	AgentConf.CollectConfSlice = append(AgentConf.CollectConfSlice, cc)
	return
}

func LoadConf(confType string, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		logs.Error("new config failed, err: %s", err)
		return
	}
	AgentConf = &Config{}
	AgentConf.LogLevel = conf.String("logs::log_level")
	if len(AgentConf.LogLevel) == 0 {
		AgentConf.LogLevel = "debug"
	}

	AgentConf.LogPath = conf.String("logs::log_path")
	if len(AgentConf.LogPath) == 0 {
		AgentConf.LogPath = "./logs"
	}

	AgentConf.ChanSize, err = conf.Int("logs::chan_size")
	if err != nil {
		AgentConf.ChanSize = 100
	}

	AgentConf.KafkaAddr = conf.String("logs::kafka_addr")
	if len(AgentConf.KafkaAddr) == 0 {
		logs.Error("kafka addr is nil")
		return
	}

	err = loadCollectConf(conf)
	if err != nil {
		logs.Error("load collect conf failed, err: %s", err)
		return
	}
	return
}


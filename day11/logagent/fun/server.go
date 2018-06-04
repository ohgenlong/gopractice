package fun

import (
	"gopractice/day11/logagent/g"
	"gopractice/day11/logagent/tailf"
	"github.com/astaxie/beego/logs"
	"time"
)

func ServerRun() {
	InitKafka(g.AgentConf.KafkaAddr)
	for {
		msg := tailf.GetOneLine()
		err := sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err: %v", err)
			time.Sleep(time.Second)
			continue
		}
 	}
}


func sendToKafka(msg *tailf.TextMsg) (err error) {
	err = SendMsgToKafka(msg.Msg, msg.Topic)
	// logs.Debug("send to kafka msg: %v, topic: %v", msg.Msg, msg.Topic)
	return
}
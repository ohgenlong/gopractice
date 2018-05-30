package fun


import (
	"gopractice/day11/logagent/tailf"
	"github.com/astaxie/beego/logs"
	"time"
)

func ServerRun() (err error) {

	for {
		msg := tailf.GetOneLine()
		err := sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err: %v", err)
			time.Sleep(time.Second)
			continue
		}
 	}

	return
}


func sendToKafka(msg *tailf.TextMsg) (err error) {

	logs.Debug("send to kafka msg: %v, topic: %v", msg.Msg, msg.Topic)
	return
}
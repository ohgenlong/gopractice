package fun


import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"time"
)


var (
	client sarama.SyncProducer
)

func InitKafka(addr string) (err error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		fmt.Println("init kafka failed: err:", err)
		return
	}
	return
}


func SendMsgToKafka(data , topic string) (err error) {

	
		msg := &sarama.ProducerMessage{}
		msg.Topic = topic
		msg.Value = sarama.StringEncoder(data)

		pid, offset, err := client.SendMessage(msg)

		if err != nil {
			fmt.Println("send msg failed: ", err)
			return err
		}
		logs.Debug("send succ, pid:%v, offset:%v, topic:%v", pid, offset, topic)
		return
		// fmt.Printf("pid: %v, offset:%v\n", pid, offset)
		// time.Sleep(10*time.Millisecond)
	
	
}


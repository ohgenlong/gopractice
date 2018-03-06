package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"221.229.164.26:9092"}, config)
	if err != nil {
		fmt.Println("produer close: err:", err)
		return
	}
	defer client.Close()

	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	msg.Value = sarama.StringEncoder("this is a good test, msg is good")

	pid, offset, err := client.SendMessage(msg)

	if err != nil {
		fmt.Println("send msg failed: ", err)
		return
	}

	fmt.Printf("pid: %v, offset:%v\n", pid, offset)

}

package main

import (
	"context"
	"fmt"

	"strconv"

	"time"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
)

const (
	Endpoint  = "127.0.0.1:8081"
	Topic     = "TestTopic"
	GroupName = "TestGroup"
)

func main() {
	producer, err := golang.NewProducer(&golang.Config{
		Endpoint:      Endpoint,
		ConsumerGroup: GroupName,
		Credentials:   &credentials.SessionCredentials{},
	}, golang.WithTopics(Topic))

	if err != nil {
		fmt.Println("NewProducer failed", err)
		panic(err)
	}

	err = producer.Start()
	if err != nil {
		fmt.Println("Start failed", err)
		panic(err)
	}

	for i := 0; i < 10; i++ {
		msg := &golang.Message{
			Topic: Topic,
			Body:  []byte("Go RocketMQ 生产者发送消息 : " + strconv.Itoa(i)),
		}

		msg.SetKeys("k1", "k2")
		msg.SetTag("test1")
		// msg.SetMessageGroup(GroupName)

		_, err = producer.Send(context.TODO(), msg)
		if err != nil {
			fmt.Println("Send failed", err)
			panic(err)
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println("Send success")

}

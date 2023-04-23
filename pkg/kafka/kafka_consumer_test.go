package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"testing"
)

func consumer() {
	// 本地搭建 kafka 服务，在 /etc/hosts 里面增加映射
	// 127.0.0.1 kafka1
	broker := "kafka1:9092"
	group := "c1"
	topics := []string{"test"}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     broker,
		"broker.address.family": "v4",
		"group.id":              group,
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest",
		"enable.auto.commit":    true,
		//"api.version.request":   false,
	})
	if err != nil {
		fmt.Println("consumer error", err.Error())
		return
	}

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		fmt.Println("sub error:", err.Error())
		return
	}

	for {
		ev := c.Poll(100)
		if ev == nil {
			continue
		}
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Println("get message:", e.Key, string(e.Value), e.TopicPartition)

		case kafka.Error:
			fmt.Println("kafka error: ", e.Error())
		default:
			fmt.Println("ignore")
		}

	}
}

func TestConsumer(t *testing.T) {
	consumer()
}

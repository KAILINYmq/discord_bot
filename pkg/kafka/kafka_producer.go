package kafka

import (
	"DiscordRolesBot/global"
	"DiscordRolesBot/pkg/transpoint"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"sync"
)

var producerPool = sync.Pool{
	New: func() interface{} {
		return NewProducer()
	},
}

func NewProducer() *kafka.Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": global.Config.Kafka.Broker,
		"acks":              1,
		"retries":           3,
		//"enable.idempotence": true,
		//"auto.create.topics.enable": true,
	})
	if err != nil {
		fmt.Println("init kafka produce error:", err.Error())
		return nil
	}
	return producer
}

func ReportData(topic string, c <-chan transpoint.MetricReportInfo) {

	for data := range c {
		go func(data transpoint.MetricReportInfo) {
			fmt.Println("receive report data:", data)
			report(topic, data)
		}(data)
	}

}

func report(topic string, data transpoint.MetricReportInfo) {
	producer := producerPool.Get().(*kafka.Producer)
	byteData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("kafka produce err:", err.Error())
	}
	if producer != nil {

		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: byteData,
		}, nil)
		if err != nil {

		}
		producerPool.Put(producer)
	} else {
		fmt.Println("producer is null")
	}
}

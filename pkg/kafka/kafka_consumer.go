package kafka

import (
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type kafkaConf struct {
	broker     string // bootstrap.servers
	family     string // broker.address.family
	group      string // group.id
	timeout    int    // session.timeout.ms
	reset      string // auto.offset.reset
	autoCommit bool   // enable.auto.commit
}

type OptionFun func(opt *kafkaConf)

type kafkaJob struct {
	consumer *kafka.Consumer
	topic    []string
	funcs    map[string]DoFunc
}

func SetBroker(broker string) OptionFun {
	return func(opt *kafkaConf) {
		opt.broker = broker
	}
}

func SetGroup(group string) OptionFun {
	return func(opt *kafkaConf) {
		opt.group = group
	}
}

type funcJob func()

func (fun funcJob) Run() {
	fun()
}

type DoFunc func(message *kafka.Message)

func NewKafka(opts ...OptionFun) (*kafkaJob, error) {

	kc := kafkaConf{
		broker:     "",
		family:     "v4",
		group:      "default",
		timeout:    6000,
		reset:      "latest",
		autoCommit: true,
	}

	for _, fun := range opts {
		fun(&kc)
	}
	if kc.broker == "" {
		return nil, errors.New("broker is empty")
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":       kc.broker,
		"broker.address.family":   kc.family,
		"group.id":                kc.group,
		"session.timeout.ms":      kc.timeout,
		"auto.offset.reset":       kc.reset,
		"enable.auto.commit":      kc.autoCommit,
		"auto.commit.interval.ms": 2000, // auto commit 时间 2 秒
		"max.poll.interval.ms":    600000,
		"heartbeat.interval.ms":   2000,
	})
	if err != nil {
		fmt.Println("kafka consumer create error", err.Error())
		return nil, err
	}

	kj := kafkaJob{
		consumer: c,
		topic:    make([]string, 0),
		funcs:    make(map[string]DoFunc, 0),
	}

	return &kj, nil
}

func (kj *kafkaJob) ConsumerRun() {

	err := kj.consumer.SubscribeTopics(kj.topic, nil)
	if err != nil {
		fmt.Println("sub error:", err.Error())
		return
	}

	for {
		ev := kj.consumer.Poll(100)
		if ev == nil {
			continue
		}
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Println(e.TopicPartition)
			if f, ok := kj.funcs[*e.TopicPartition.Topic]; ok {
				// 消费消息后异步调用，没法保证顺序，同步调用性能差点，但是可以保证顺序 😄
				//go f(e)
				f(e)
			}
			//tp, err := kj.consumer.Commit()
			//if err != nil {
			//	global.Logger.Info(fmt.Sprintf("topic partition: %+v", tp))
			//}
		case kafka.Error:
			fmt.Println("kafka error: ", e.Error())
		default:
			fmt.Println("ignore")
		}
	}
}

func (kj *kafkaJob) RegisterTopic(topic string, f DoFunc) error {
	if _, ok := kj.funcs[topic]; ok {
		return fmt.Errorf("topic: %v has already rehister", topic)
	}
	if kj.topic == nil {
		kj.topic = make([]string, 0)
	}
	kj.topic = append(kj.topic, topic)

	if kj.funcs == nil {
		kj.funcs = make(map[string]DoFunc, 0)
	}
	kj.funcs[topic] = f
	return nil
}

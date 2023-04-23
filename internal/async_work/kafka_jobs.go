package async_work

import (
	"DiscordRolesBot/global"
	"DiscordRolesBot/pkg/kafka"
	"DiscordRolesBot/pkg/transpoint"
	"fmt"
	"os"
)

var (
	MetricReportChan = make(chan transpoint.MetricReportInfo, 100)
	enableReportData = false
	reportTopic      string
)

func init() {
	enableReport := os.Getenv("ENABLE_REPORT")
	reportTopic = os.Getenv("KAFKA_REPORT_TOPIC")
	if enableReport == "true" && reportTopic != "" {
		enableReportData = true
	}
}

// kafka 消费
func Jobs() {
	job, err := kafka.NewKafka(kafka.SetBroker(global.Config.Kafka.Broker), kafka.SetGroup(global.Config.Kafka.Group))
	if err != nil {
		errMsg := fmt.Sprintf("kafka create consumer error: %v", err.Error())
		global.Logger.Error(errMsg)
		return
	}
	job.RegisterTopic(global.Config.Kafka.EventTopic, eventTopic)

	go func() {
		job.ConsumerRun()
	}()

	if enableReportData {
		go func() {
			kafka.ReportData(reportTopic, MetricReportChan)
		}()
	}
}

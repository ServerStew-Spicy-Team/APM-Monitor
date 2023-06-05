package reporter

import (
	"APM-Monitor/internal/apm-monitor/collector"
	"APM-Monitor/internal/pkg/log"
	"APM-Monitor/pkg/kafka"
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func Report(c collector.Collector) {
	val := <-c.CollectorReturnMetric()
	if val.Length() == 0 {
		log.Errorw("metrics are nil")
		return
	}
	//str, err := val.Encode()
	ip := viper.GetString("ip")
	msg := &sarama.ProducerMessage{
		Topic: c.GetTopic(),
		Key:   sarama.StringEncoder(ip),
		Value: val,
	}
	_, _, err := kafka.Pro().P.SendMessage(msg)
	if err != nil {
		log.Errorw("Failed to send message:", err)
	}
	//log.Infow("Message sent to", "partition", partition, "at offset", offset)
}

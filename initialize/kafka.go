package initialize

import (
	"dancin-api/global"
	"github.com/segmentio/kafka-go"
	"strings"
)

func KafkaWriter() {
	global.KAFKA_WRITER = &kafka.Writer{
		Addr:     kafka.TCP(global.CONFIG.Report.Path),
		Topic:    global.CONFIG.Report.Topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func KafkaReader() {
	brokers := strings.Split(global.CONFIG.Report.Path, ",")
	global.KAFKA_READER = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  global.CONFIG.Report.Group,
		Topic:    global.CONFIG.Report.Topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

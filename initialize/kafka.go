package initialize

import (
	"danci-api/global"
	kafka "github.com/segmentio/kafka-go"
	"strings"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// 启动一个 kafka 生产者
func KafkaWriter() {
	kafkaURL := "localhost:9092"
	topic := "maxSB"
	global.GVA_KAFKA_WRITER = newKafkaWriter(kafkaURL, topic)
}

func KafkaReader() {
	kafkaURL := "localhost:9092"
	topic := "maxSB"
	global.GVA_KAFKA = getKafkaReader(kafkaURL, topic, "maxSB")
}

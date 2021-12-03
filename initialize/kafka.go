package initialize

import (
	"context"
	"dancin-api/global"
	"github.com/segmentio/kafka-go"
)

func KafkaConn() error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", global.CONFIG.Report.Path, global.CONFIG.Report.Topic, 0)
	if err != nil {
		global.KAFKA = nil
		return err
	}
	global.KAFKA = conn
	return nil
}

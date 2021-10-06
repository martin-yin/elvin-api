package global

import (
	"danci-api/config"
	"go.uber.org/zap"

	"github.com/go-redis/redis"
	kafka "github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB           *gorm.DB
	GVA_REDIS        *redis.Client
	GVA_CONFIG       config.Server
	GVA_VP           *viper.Viper
	GVA_LOG          *zap.Logger
	GVA_KAFKA_WRITER *kafka.Writer
	GVA_KAFKA        *kafka.Reader
)

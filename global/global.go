package global

import (
	"dancin-api/config"
	"go.uber.org/zap"

	"github.com/go-redis/redis"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GORMDB *gorm.DB
	REDIS  *redis.Client
	CONFIG config.Server
	VIPER  *viper.Viper
	LOGGER *zap.Logger
	KAFKA  *kafka.Conn
)
